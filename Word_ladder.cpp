//BFS solution, this solution is not effiective
class Solution {
    public:
        int ladderLength(string start, string end, unordered_set<string> &dict) {
           if(dict.empty())
                return 0;
           
           queue<string>q;
           q.push(start);
           unordered_map<string, int> visited;
           unordered_set<string> unvisited = dict;
           visited[start] = 1;
           unvisited.erase(start);
           
           while(!q.empty())
           {
               string word = q.front();
               q.pop();
               auto iter = unvisited.begin();
               while(iter != unvisited.end())
               {
                   string adjword = *iter;
                   if(onediff(word, adjword))
                   {
                        visited[adjword] = visited[word] + 1;
                        if(adjword == end)
                            return visited[adjword];
                        iter = unvisited.erase(iter);
                        q.push(adjword);
                   }
                   else
                        iter++;
               }
           }
            return 0;
        }
        
        inline bool onediff(string &str1, string &str2)
        {
            int count = 0;
            for(int i = 0; i < str1.size(); i++)
            {
             if(str1[i] != str2[i])
             {
                count++;
             }
              if(count > 1)
                    return false;
            }
            return count == 1;
        }
    };  

//DBFS solution, this is a effective solution

 /*DBFS method, each time generate all reachable words in the dict that can be transformed from start in the ith steps, and delete this word in dict. 
Terminate until reach the same word in the other frontier and return dist+1, or a queue is empty and the path does not exist.
    */

    class Solution {
    public:
        public:
            int ladderLength(string start, string end, unordered_set<string> &dict) {
                queue<string> q1, q2, q3;                        // q1, q2 are the queues for two ends
                unordered_set<string>   frontier1, frontier2;       //frontier1, frontier2 store words that were just reached on both ends
                int dist = 1;
                q1.push(start); q2.push(end);
                frontier1.insert(start); frontier2.insert(end);

                while ( !q1.empty() && !q2.empty()) {       //If one queue is empty, no connection can be built

                    if(q1.size() <= q2.size()){
                        frontier1.clear();
                        while(!q1.empty()){
                            string word = q1.front(); q1.pop();
                            for (int i = 0; i < word.size(); i++) {
                                string temp = word;
                                for (char c = 'a'; c <= 'z'; c++) {
                                    if(c==word[i])  continue;
                                    temp[i] =c;

                                    if(frontier2.count(temp) > 0)     return dist+1;

                                    if (dict.count(temp) > 0) {
                                    //use q3 to store the intermediate words that can be reached in the next step
                                        q3.push(temp);         
                                        dict.erase(temp);
                                        frontier1.insert(temp);
                                    }
                                }
                            }
                        }
                        dist++;
                        swap(q1, q3);
                    }
                    else{
                        frontier2.clear();
                        while(!q2.empty()){
                            string word = q2.front(); q2.pop();
                            for (int i = 0; i < word.size(); i++) {
                                string temp = word;
                                for (char c = 'a'; c <= 'z'; c++) {
                                    if(c==word[i])  continue;
                                    temp[i] =c;

                                    if(frontier1.count(temp) > 0)     return dist+1;

                                    if (dict.count(temp) > 0) {
                                        //use q3 to store the intermediate words that can be reached in the next step
                                        q3.push(temp);          
                                        dict.erase(temp);
                                        frontier2.insert(temp);
                                    }
                                }
                            }
                        }
                        dist++;
                        swap(q2, q3);
                    }

                }
                return 0;
            }
    };  