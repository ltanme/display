# 1. display
A project using GoFrame,this is a system designed for displaying digital content in households, which includes features such as posting notifications, countdowns, and showing stroke orders of Chinese characters.

# 2. Quick start

## 2.1 If you use Sqlite

1. Download the source code
```
git clone https://github.com/ltanme/display.git
cd display
```

2. Run display
```
cp manifest/config/config.example.yaml manifest/config/config.yaml
gf run main.go
```

3. Enjoy it
then open http://127.0.0.1:81/ and enjoy it.

> user: goframe
> password: 123456

## 2.2 If you use mysql 

1. Download the source code
```
git clone https://github.com/ltanme/display.git
cd display
```

2. Update config
copy manifest/config/config.yaml and edit database.default config
```
cp manifest/config/config.example.yaml manifest/config/config.yaml
```

3. Import Db
Import manifest/document/focus.sql to your Mysql

4. Run display And Enjoy
```
gf run main.go
```

then open http://127.0.0.1:81/ and enjoy it.

> user: goframe
> password: 123456
