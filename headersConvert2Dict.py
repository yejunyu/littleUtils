s = '''
accept:text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8
accept-encoding:gzip, deflate, br
accept-language:zh-CN,zh;q=0.9
cache-control:no-cache
cookie:gitbook:sess=eyJ2YXJpYW50IjoxLCJhbm9ueW1vdXNJZCI6IjdkODA4MjI3LWJiMzQtNGI5MS05MDRjLTczYTRmZGU0NTA5ZiIsImNzcmZTZWNyZXQiOiJ6SWJKV3g0XzVabW50VjZUMlp6c2ZvLXIifQ==; gitbook:sess.sig=TzpnjvP9zImm6azI3ipk3d-M74Y; _ga=GA1.2.1206922231.1517479696; _gid=GA1.2.1155033187.1517479696
dnt:1
pragma:no-cache
upgrade-insecure-requests:1
user-agent:Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36
'''

s = s.strip().split('\n')
s = {x.split(':', 1)[0] : x.split(':', 1)[1] for x in s}
print (s)
