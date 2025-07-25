# Typora Theme


{{< admonition type=abstract title="导语" open=true >}}
分享一个好看的 Typora 主题
{{< /admonition >}}

<!--more-->

## 效果

本 Typora 软件主题是在 [`Purple`](https://theme.typora.io/theme/Purple/) 之上修改而来，具体效果如下：

{{< figure src="/posts/typora-theme/theme.png" title="" >}}

## Source

由于实现源码太长，给出下载连接：[kyden.css](https://github.com/kydenul/dotfiles/blob/master/kyden.css)

```css
@include-when-export url(https://fonts.loli.net/css?family=Open+Sans:400italic,700italic,700,400&subset=latin,latin-ext);

/* open-sans-regular - latin-ext_latin */

@font-face {
    font-family: "Open Sans";
    font-style: normal;
    font-weight: normal;
    src: local("Open Sans Regular"), local("OpenSans-Regular"), url("./github/open-sans-v17-latin-ext_latin-regular.woff2") format("woff2");
    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD, U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}


/* open-sans-italic - latin-ext_latin */

@font-face {
    font-family: "Open Sans";
    font-style: italic;
    font-weight: normal;
    src: local("Open Sans Italic"), local("OpenSans-Italic"), url("./github/open-sans-v17-latin-ext_latin-italic.woff2") format("woff2");
    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD, U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}


/* open-sans-700 - latin-ext_latin */

@font-face {
    font-family: "Open Sans";
    font-style: normal;
    font-weight: bold;
    src: local("Open Sans Bold"), local("OpenSans-Bold"), url("./github/open-sans-v17-latin-ext_latin-700.woff2") format("woff2");
    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD, U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}


/* open-sans-700italic - latin-ext_latin */

@font-face {
    font-family: "Open Sans";
    font-style: italic;
    font-weight: bold;
    src: local("Open Sans Bold Italic"), local("OpenSans-BoldItalic"), url("./github/open-sans-v17-latin-ext_latin-700italic.woff2") format("woff2");
    unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD, U+0100-024F, U+0259, U+1E00-1EFF, U+2020, U+20A0-20AB, U+20AD-20CF, U+2113, U+2C60-2C7F, U+A720-A7FF;
}

:root {
    --title-color: #8064a9;
    --text-color: #444444;
    --light-text-color: #666666;
    --lighter-text-color: #888888;
    /* --link-color: #2aa899; */
    /* --code-color: #745fb5; */

    --link-color: #745fb5;
    /* --code-color: #2aa899; */
    --code-color: #ec71b7;

    --shadow-color: #eee;
    --border-quote: rgba(116, 95, 181, 0.2);
    --border-quote-grey: #c8c8c8;
    --border: #e7e7e7;
    --link-bottom: #bbb;
    --shadow: 3px 3px 10px var(--shadow-color);
    --inline-code-bg: #f4f2f9;
    --header-weight: normal;
    --side-bar-bg-color: #fafafa;
    --control-text-color: var(var(--light-text-color));
    --active-file-text-color: var(--title-color);
    --active-file-bg-color: var(--shadow-color);
    --item-hover-bg-color: var(--shadow-color);
    --active-file-border-color: var(var(--title-color));
    --base-font: "Open Sans", "Clear Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
    --title-font: "EB Garamond", "Source Sans Pro", serif;
    --monospace: Courier, Monospace !important;
}


/* 打印 */

@media print {
    html {
        font-size: 0.9rem;
    }

    table,
    pre {
        page-break-inside: avoid;
    }

    pre {
        word-wrap: break-word;
    }

    #write {
        max-width: 100%;
    }

    @page {
        size: A2;
        font-size: 0.2rem;
        /* PDF output size */
        margin-left: 0;
        margin-right: 0;
    }
}

html {
    font-size: 16px;
    -webkit-text-size-adjust: 100%;
    -ms-text-size-adjust: 100%;
    text-rendering: optimizelegibility;
    -webkit-font-smoothing: initial;
}

body {
    color: var(--text-color);
    -webkit-font-smoothing: antialiased;
    line-height: 1.6;
    letter-spacing: 0;
    overflow-x: hidden;
}


/* 页边距 和 页面大小 */

#write {
    font-family: var(--base-font);
    /* max-width: 914px; */
    margin: 0 auto;
    padding: 1rem 4rem;
    padding-bottom: 100px;
}

#write p {
    line-height: 1.6rem;
    word-spacing: 0.05rem;
}

body>*:first-child {
    margin-top: 0 !important;
}

body>*:last-child {
    margin-bottom: 0 !important;
}


/* Link 链接 */

a {
    color: var(--link-color);
    text-decoration: none;
}

#write a {
    border-bottom: 1px solid var(--link-bottom);
}

.md-content {
    color: var(--light-text-color);
}

h1,
h2,
h3,
h4,
h5,
h6 {
    position: relative;
    margin-top: 2rem;
    margin-bottom: 1rem;
    font-weight: var(--header-weight);
    line-height: 1.3;
    cursor: text;
    color: var(--title-color);
    font-family: var(--title-font);
}

h1 {
    text-align: center;
    margin-bottom: 2rem;
    line-height: 1.5;
}

h1:after {
    content: "";
    display: block;
    margin: 0.2rem auto 0;
    width: 6rem;
    height: 2px;
    border-bottom: 2px solid var(--title-color);
}

h2 {
    padding-left: 0.4em;
    border-left: 0.4em solid var(--title-color);
    border-bottom: 1px solid var(--title-color);
}

h3 {
    padding-left: 0.2em;
    border-left: 0.2em dashed #2aa899;
}

h1 tt,
h1 code {
    font-size: inherit;
}

h2 tt,
h2 code {
    font-size: inherit;
}

h3 tt,
h3 code {
    font-size: inherit;
}

h4 tt,
h4 code {
    font-size: inherit;
}

h5 tt,
h5 code {
    font-size: inherit;
}

h6 tt,
h6 code {
    font-size: inherit;
}

p,
blockquote,
ul,
ol,
dl,
table {
    margin: 0.8em 0;
}


/* horizontal rule */

hr {
    margin: 1.5em auto;
    border-top: 1px solid var(--border);
}


/* 列表 */

li>ol,
li>ul {
    margin: 0 0;
}

li p.first {
    display: inline-block;
}

ul,
ol {
    padding-left: 2rem;
}

ul:first-child,
ol:first-child {
    margin-top: 0;
}

ul:last-child,
ol:last-child {
    margin-bottom: 0;
}

#write ol li,
ul li {
    padding-left: 0.1rem;
}


/* 引用 */

blockquote {
    border-left: 0.3rem solid var(--border-quote);
    padding-left: 1em;
    font-family: var(--base-font);
}


/* 表格 */

table {
    margin-bottom: 1.25rem;
}

table th,
table td {
    padding: 8px;
    line-height: 1.25rem;
    vertical-align: middle;
    border: 1px solid #ddd;
}

table th {
    font-weight: bold;
}

table thead th {
    vertical-align: middle;
}

table tr:nth-child(2n),
thead {
    background-color: #fcfcfc;
}


/* 粗体 */

#write strong {
    padding: 0 2px;
    font-weight: bold;
}


/* inline code */
code,
tt {
    padding: 2px 4px;
    border-radius: 0.3rem;
    font-family: var(--monospace);
    font-size: 0.9rem;
    color: var(--code-color);
    background-color: var(--inline-code-bg);
    margin: 0 2px;
}

#write .md-footnote {
    color: var(--code-color);
    background-color: var(--inline-code-bg);
}


/* highlight. */

#write mark {
    background: rgb(255, 255, 196);
    color: var(--text-color);
}

#write del {
    padding: 1px 2px;
}

.md-task-list-item>input {
    margin-left: -1.3em;
}

#write pre.md-meta-block {
    padding: 1rem;
    font-size: 85%;
    line-height: 1.45;
    background-color: #f7f7f7;
    border: 0;
    border-radius: 3px;
    color: #777777;
    margin-top: 0 !important;
}

.mathjax-block>.code-tooltip {
    bottom: 0.375rem;
}


/* 图片 */

.md-image>.md-meta {
    border-radius: 3px;
    font-family: var(--monospace);
    padding: 2px 0 0 4px;
    font-size: 0.9em;
    color: inherit;
}


/* 图片靠左显示 */


/* p .md-image:only-child {
  width: auto;
  text-align: left;
  margin-left: 2rem;
} */


/* 写![shadow-...]() 显示图片阴影 */

img[alt|="shadow"] {
    box-shadow: var(--shadow);
}


/* TOC */

#write a.md-toc-inner {
    line-height: 1.6;
    white-space: pre-line;
    border-bottom: none;
    color: var(--light-text-color);
    font-size: 0.9rem;
}

.md-toc-h1 .md-toc-inner {
    margin-left: 0;
    font-weight: normal;
}

header,
.context-menu,
.megamenu-content,
footer {
    font-family: var(--base-font);
}

.file-node-content:hover .file-node-icon,
.file-node-content:hover .file-node-open-state {
    visibility: visible;
}

.md-lang {
    color: #b4654d;
}

.html-for-mac .context-menu {
    --item-hover-bg-color: #e6f0fe;
}


/* Code fences */


/* border, bg, font */

.md-fences {
    border: 1px solid var(--border);
    border-radius: 5px;
    background: #fdfdfd !important;
    font-size: 0.9rem;
}


/* 代码框阴影 */


/* #write pre.md-fences {
  display: block;
  -webkit-overflow-scrolling: touch;
  box-shadow: var(--shadow);
} */

.cm-s-inner .cm-variable {
    color: var(--text-color);
}

.cm-s-inner {
    padding: 0.25rem;
    border-radius: 0.25rem;
}

.cm-s-inner.CodeMirror,
.cm-s-inner .CodeMirror-gutters {
    color: #3a3432 !important;
    border: none;
}

.cm-s-inner .CodeMirror-gutters {
    color: #6d8a88;
}

.cm-s-inner .CodeMirror-linenumber {
    padding: 0 0.1rem 0 0.3rem;
    color: #b8b5b4;
}

.cm-s-inner .CodeMirror-matchingbracket {
    text-decoration: underline;
    color: #a34e8f !important;
}

#fences-auto-suggest .active {
    background: #ddd;
}

.cm-s-inner span.cm-comment {
    color: #9daab6;
}

.cm-s-inner span.cm-builtin {
    color: #0451a5;
}


/* language tip */

#write .code-tooltip {
    border: 1px solid var(--border);
}

.auto-suggest-container {
    border-color: #b4b4b4;
}

.auto-suggest-container .autoComplt-hint.active {
    background: #b4b4b4;
    color: inherit;
}


/* task list */

#write .md-task-list-item>input {
    -webkit-appearance: initial;
    display: block;
    position: absolute;
    border: 1px solid #b4b4b4;
    border-radius: 0.2rem;
    margin-top: 0.3rem;
    height: 1rem;
    width: 1rem;
    transition: background 0.3s;
}

#write .md-task-list-item>input:focus {
    outline: none;
    box-shadow: none;
}

#write .md-task-list-item>input:hover {
    background: #ddd;
}

#write .md-task-list-item>input[checked]::before {
    content: "";
    position: absolute;
    top: 20%;
    left: 50%;
    height: 60%;
    width: 2px;
    transform: rotate(40deg);
    background: #333;
}

#write .md-task-list-item>input[checked]::after {
    content: "";
    position: absolute;
    top: 46%;
    left: 25%;
    height: 30%;
    width: 2px;
    transform: rotate(-40deg);
    background: #333;
}

#write .md-task-list-item>p {
    transition: color 0.3s, opacity 0.3s;
}

#write .md-task-list-item.task-list-done>p {
    color: #b4b4b4;
    text-decoration: line-through;
}

#write .md-task-list-item.task-list-done>p>.md-emoji {
    opacity: 0.5;
}

#write .md-task-list-item.task-list-done>p>.md-link>a {
    opacity: 0.6;
}


/* sidebar */

#typora-sidebar,
.typora-node #typora-sidebar {
    box-shadow: 3px 0px 10px var(--shadow-color);
}

.sidebar-content-content {
    font-size: 0.9rem;
}
```

## Reference

- [Typora](https://typora.io/)
- [Purple](https://theme.typora.io/theme/Purple/)


---

> Author: [kyden](https://github.com/kydenul)  
> URL: http://kydenul.github.io/posts/typora-theme/  

