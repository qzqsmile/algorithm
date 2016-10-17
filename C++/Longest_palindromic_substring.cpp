/*
����㷨Ҫע�����¼���
1.��ĳ���ַ�Ϊ������������չ�Ӷ�ȷ���Ƿ��ǻ���
2.ȷ��ʱҪע����ż�������⣬��Ϊż���������������ȣ����Կ���ͨ���ж��Ƿ����ȥ����������
3.ĳһ�����������ֽ����ڸ��ʼ���ַ����ܴ������Ļ��� �� *aaaaab*��������Ҫ�жϵ�һ��a��ԭ����������i�����ǰ���
 Ѱ��begin �� end����бȽϣ���ȥ��һa֮���ڵ�bһ����ͣ��������Ҳ���±ߴ�����i = end��ԭ��
*/
class Solution {
public:
    string longestPalindrome(string s) {
        if (s.empty()) return "";
        if (s.size() == 1) return s;
        int max_len = 0, max_begin = 0;
        
        for(int i = 0; i < s.size();)
        {
            if (s.size() - i <= max_len / 2) break;
            int begin,end;
            end = i;
            while((end < (s.size()-1)) && (s[end] == s[end+1])) end++;
            end++;
            begin = i -1;
            i = end;
            while((end < s.size()) && (begin >= 0) && (s[begin] == s[end]))
            {
                begin--;
                end++;
            }
            int new_len = end-begin-1;
            if(new_len > max_len)
            {
                max_len = new_len;
                max_begin = begin+1;
            }
        }
        
        return s.substr(max_begin, max_len);
    }
};