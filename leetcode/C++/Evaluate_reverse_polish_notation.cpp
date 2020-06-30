class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> s;
        if(!tokens.size())
            return 0;
        for(int i = 0; i < tokens.size(); i++)
        {
            if((tokens[i] == "*") || (tokens[i] == "+") || (tokens[i] == "-") || (tokens[i] == "/"))
            {
                int a = s.top();
                s.pop();
                int b = s.top();
                s.pop();
                if(tokens[i] == "*")
                    s.push(a*b);
                if(tokens[i] == "/")
                    s.push(b/a);
                if(tokens[i] == "-")
                    s.push(b-a);
                if(tokens[i] == "+")
                    s.push(a+b);                
            }
            else
            {
                int temp;
                stringstream s1;
                s1 << tokens[i];
                s1 >> temp;
                s.push(temp);
            }
        }
        
        return s.top();
    }
};