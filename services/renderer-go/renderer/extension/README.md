## collapse Extention
goldmarkの折りたたみ記法（独自）をサポートするための拡張機能です

## usage

`|>(`で始まり，その直後に続く文字が要約部分，次の行からは折りたたまれる内容になります

終了時は，`)`のみの行で閉じて終わります


## testcase

入力
```
|>(サンプルの折りたたみ
これは折りたたまれるコンテンツです。
- hoge
    - huga
)
```

出力
```
<details class="collapse">
<summary>サンプルの折りたたみ</summary>
<p>これは折りたたまれるコンテンツです。</p>
<ul>
<li>hoge
<ul>
<li>huga</li>
</ul>
</li>
</ul>
</details>
```