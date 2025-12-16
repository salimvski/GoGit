package api

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
    "fmt"
	"gogit/internal/models"
)

type CacheEntry struct {
    CachedAt   time.Time   `json:"cached_at"`
    TTLSeconds int         `json:"ttl_seconds"`
    Type       string      `json:"type"`
    Data       interface{} `json:"data"`
}

type Cache struct {
    Version  int                   `json:"version"`
    Metadata CacheMetadata         `json:"metadata"`
    Entries  map[string]CacheEntry `json:"entries"`
}

type CacheMetadata struct {
    CreatedAt   time.Time `json:"created_at"`
    EntriesCount int      `json:"entries_count"`
}

var cache *Cache

func InitCache() error {
    cache = &Cache{
        Version: 1,
        Metadata: CacheMetadata{
            CreatedAt:   time.Now(),
            EntriesCount: 0,
        },
        Entries: make(map[string]CacheEntry),
    }
    return loadCacheFromDisk()
}

func loadCacheFromDisk() error {
    cacheFile := filepath.Join(getCacheDir(), "cache.json")
    
    data, err := os.ReadFile(cacheFile)
    if err != nil {
        if os.IsNotExist(err) {
            return nil
        }
        return err
    }
    
    return json.Unmarshal(data, cache)
}

func getCacheDir() string {
    cacheDir, err := os.UserCacheDir()
    if err != nil {
        cacheDir = filepath.Join(os.Getenv("HOME"), ".cache")
    }
    gogitCache := filepath.Join(cacheDir, "gogit")
    os.MkdirAll(gogitCache, 0755)
    return gogitCache
}

func getCachedUser(username string) (*models.User, error) {
    key := "user:" + username
    
    entry, exists := cache.Entries[key]
    if !exists {
        return nil, errors.New("not in cache")
    }
    
    if time.Now().Sub(entry.CachedAt) > time.Duration(entry.TTLSeconds)*time.Second {
        return nil, errors.New("cache expired")
    }
    
    dataBytes, err := json.Marshal(entry.Data)
    if err != nil {
        return nil, fmt.Errorf("marshaling cache data: %w", err)
    }
    
    var user models.User
    if err := json.Unmarshal(dataBytes, &user); err != nil {
        return nil, fmt.Errorf("unmarshaling cache data: %w", err)
    }
    
    return &user, nil
}

func cacheUser(username string, user *models.User, ttlSeconds int) error {
    key := "user:" + username
    
    entry := CacheEntry{
        CachedAt:   time.Now(),
        TTLSeconds: ttlSeconds,
        Type:       "user",
        Data:       user,
    }
    
    cache.Entries[key] = entry
    cache.Metadata.EntriesCount = len(cache.Entries)
    
    return saveCacheToDisk()
}

func saveCacheToDisk() error {
    jsonData, err := json.MarshalIndent(cache, "", "  ")
    if err != nil {
        return err
    }
    
    cacheFile := filepath.Join(getCacheDir(), "cache.json")
    return os.WriteFile(cacheFile, jsonData, 0644)
}