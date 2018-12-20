s = '''
VISITOR_INFO1_LIVE=MzngxulmUNg; CONSENT=YES+US.zh-CN+; s_gl=52472a9550e56850cd0bf4e212c0b890cwIAAABJTg==; PREF=hl=en&gl=IN&fms1=10000&fms2=10000&f1=50000000&f5=30&al=zh-CN; YSC=mO9L20vPmeU; GPS=1'''
s = s.strip().split(';')
s = {x.split('=', 1)[0]: x.split('=', 1)[1] for x in s}
print(s)