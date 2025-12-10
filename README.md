# **📋 GoGit Insights - Feature List ONLY**

## **MVP Features (Week 1)**
1. **Fetch Basic Repo Info**
   - List user repositories
   - Show name, language, stars
   - Handle API errors 

2. **Concurrent Fetching**
   - Worker pool for parallel API calls
   - Progress indicator

3. **Basic Statistics**
   - Language distribution
   - Total/Average stars
   - Most popular repo

## **V1 Features (Week 2)**
4. **Caching**
   - Local file cache
   - TTL management

5. **Better Display**
   - Colored terminal output
   - ASCII progress bars
   - Simple charts

6. **Profiling**
   - pprof endpoints
   - Performance benchmarks

## **Commands**
   # Primary nouns
   - gogit repo list <user>        # list repos
   - gogit repo stats <user>       # deep repo statistics

   - gogit user view <user>        # show user profile card
   - gogit user stats <user>       # follower/following, contributions, etc.

---