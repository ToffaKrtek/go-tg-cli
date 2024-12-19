### USE

```
./go-tg-cli -msg="TEXT MSG" -file="/path/to/file" -image="/path/to/image"
# also you can:
# -mode="HTML||Markdown"
# -chat_id="123||or env"
# -topic_id="123||or env"
# -token="123||or env"
```

```
./s3-uploader -file=./project.zip \
  -url=s3.example.ru \
  -object=dir/project.zip \
  -bucket=test \
  -access=<access-token> \
  -secret=<secret-key>
# or fron .env
```
