# Goraph

- ダイクストラ法による経路検索
- OpenStreetMapからの読み込み
- 

## グラフの読み込み

### CSVファイルによる読み込み

```csv
from,to,cost
node_id,node_id,cost
```

node_idは、文字列でも数字でも、欠番があってもよい。

### プロトコルバッファによる読み込み

高速。