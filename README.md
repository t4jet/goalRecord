# GoalRecord

## backend
使い方
```docker-compose up -d --build```

データ投入
``` curl http://localhost:8080/api/tasks -X POST -H "Content-Type: application/json" -d '{"name": "掃除"}'```

ポートがふさがていたら
``` sudo service mysql stop```

## frontend
使い方

最初にする
```npm install```

起動
```npm run dev```
