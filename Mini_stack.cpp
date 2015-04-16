class MinStack {
public:
    stack<int> data;
    stack<int> min;
    void push(int x) {
        if(min.empty())
        {
            data.push(x);
            min.push(x);
            return ;
        }
        data.push(x);
        
        if(x <= min.top())
            min.push(x);
    }

    void pop() {
            if (data.top() == min.top())
                min.pop();
            data.pop();
    }

    int top() {
        return data.top();
    }

    int getMin() {
        return min.top();
    }
};