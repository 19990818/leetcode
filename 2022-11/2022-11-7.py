class Solution:
    def ambiguousCoordinates(self, s: str) -> List[str]:
        res=[]
        m=dict()
        for i in range(2,len(s)-1):
            if self.isValid(s[1:i]) and self.isValid(s[i:-1]):
                res.append(s[0:i]+', '+s[i:])
            for j in range(1,i+1):
                for k in range(i,len(s)):
                    left,right=s[1:j]+'.'+s[j:i],s[i:k]+'.'+s[k:-1]
                    temp=""
                    if self.isValid(left) and self.isValid(right):
                        temp='('+left+', '+right+')'
                    elif self.isValid(left) and self.isValid(s[i:-1]):
                        temp='('+left+', '+s[i:-1]+')'
                    elif self.isValid(right) and self.isValid(s[1:i]):
                        temp='('+s[1:i]+', '+right+')'
                    if temp!="" and m.get(temp) is None:
                        res.append(temp)
                        m[temp]=1                    
        return res
    def isValid(self,src:str)->bool:
        if src=="" or src==".":return False
        arr=src.split('.')
        if len(arr[0])==0 or (arr[0][0]=='0' and len(arr[0])>1):
            return False
        if len(arr)>1 and (len(arr[1])==0 or arr[1][-1]=='0'):
            return False
        return True