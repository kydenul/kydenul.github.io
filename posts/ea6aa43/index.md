# Claude Code Guide


{{< admonition type=abstract title="导语" open=true >}}
**这是导语部分**
{{< /admonition >}}

<!--more-->
### 一、Slash commands

Slash commands 是 Claude Code 中的斜杆指令。

### I. 内置的 Slash commands

> [!TIP]
> 个人常用的内置 Slash commands

| Command | Purpose |
| --- | --- |
| `/agents` | Manage custom AI subagents for specialized tasks <br> 管理自定义 AI 子代理以执行专业任务 |
| `/clear` | Clear conversation history <br> 清除对话历史记录 |
| `/compact [instructions]` | Compact conversation with optional focus instructions <br> 压缩对话 |
| `/context` | Visualize current context usage as a colored grid <br> 以彩色网格形式可视化当前上下文使用情况 |
| `/exit` | Exit the REPL  <br> 退出 REPL |
| `/init` | Initialize project with `CLAUDE.md` guide <br> `初始化 CLAUDE.md` |
| `/mcp` | Manage MCP server connections and OAuth authentication <br> 管理 MCP 服务器连接和 OAuth 认证 |
| `/model` | Select or change the AI model <br> 选择或更改 AI 模型 |
| `/vim` | Enter vim mode for alternating insert and command modes <br> 进入 vim 模式以切换插入模式和命令模式 |

---

### II. Custom Slash Commands

自定义 Slash Commands 允许将常用提示语保存为 Markdown 文件，以便在 Claude Code 中使用。
Slash Command 通过作用域（**Project-Specific** 或 **Personal**）进行组织，并通过目录结构支持命名空间。

>[!TIP] Syntax
> `/<command-name> [arguments]`
>
> - **`command-name`**: Name derived from the Markdown filename (without `.md` extension)
> - **`[arguments]`**: Optional arguments passed to the command

#### 1. Project-specific Slash Commands

Project-specific Slash Commands 存储于项目中的 `.claude/commands/` 文件夹中。

#### 2. Personal Slash Commands

Personal Slash Commands 存储于用户文件夹中的 `~/.claude/commands/` 文件夹中。

#### 3. Namespacing

Namespacing 支持使用 `子目录` 来分组相关的命令 => 子目录会出现在命令描述中，但不影响命令名称。

> [!TIP]
>
> - 当 Project-specific 命令与 Personal 命令名称相同时，Proect-specific 具有优先权，Personal 命令将被默认忽略。
> - 不同子目录中的命令可以共享名称，因为可以利用命令描述中的子目录进行区分。

#### 4. Arguments

Slash Commands 支持参数，可以在命令名称后面添加参数，以便在命令执行时传递参数。

- `$ARGUMENTS`: 所有参数

    ```bash
        # Command definition
    echo 'Fix issue #$ARGUMENTS following our coding standards' > .claude/commands/fix-issue.md

    # Usage
    > /fix-issue 123 high-priority
    # $ARGUMENTS becomes: "123 high-priority"
    ```

- `$1` / `$2` / `$3` / `$4` / `$5`: 位置参数

    ```bash
    # Command definition  
    echo 'Review PR #$1 with priority $2 and assign to $3' > .claude/commands/review-pr.md

    # Usage
    > /review-pr 456 high alice
    # $1 becomes "456", $2 becomes "high", $3 becomes "alice"
    ```

### III. Slash Commands Examples

```makrdown
---
allowed-tools: Bash(git add:*), Bash(git status:*), Bash(git commit:*)
description: Create a git commit
---

## Context

- Current git status: !`git status`
- Current git diff (staged and unstaged changes): !`git diff HEAD`
- Current branch: !`git branch --show-current`
- Recent commits: !`git log --oneline -10`

## Your task

Based on the above changes, create a single git commit.
```

### 二、Sub Agents

#### I. What is a Sub Agent?

Sub Agent 是预配置的 AI 人格，Claude Code 可以将任务委派给它们。每个子代理：

- 具有特定的目的和专业领域
- 使用与主对话分离的自己的上下文窗口
- 可以配置为允许使用特定工具
- 包含指导其行为的自定义系统提示
- 当 Claude Code 遇到与 Sub Agent 专业领域相匹配的任务时，它可以将该任务委派给专门的 Sub Agent，该 Sub Agent 将独立工作并返回结果。

#### II. 优势

- 上下文保留: 每个 Sub Agent 在自己的上下文中运行，防止主对话被污染，并使其专注于高级目标。
- 专业化专业知识: Sub Agent 可以使用特定领域的详细说明进行微调，从而提高指定任务的成功率。
- 可重用性: 创建后，Sub Agent 可以在不同项目中使用，并与您的团队共享以实现一致的工作流。
- 灵活的权限: 每个 Sub Agent 可以具有不同的工具访问级别。

#### III. Sub Agent 配置文件

子代理存储为具有 YAML 前置内容的 Markdown 文件，位置有两个：

| 类型 | 位置 | 范围 | 优先级 |
| --- | --- | --- | --- |
|项目子代理 | `.claude/agents/` | 在当前项目中可用 | 最高 |
|用户子代理 | `~/.claude/agents/` | 在所有项目中可用 | 较低 |

> [!NOTE]
> 当子代理名称冲突时，项目级别的子代理优先于用户级别的子代理。

##### 文件格式

```Markdown
---
name: your-sub-agent-name
description: Description of when this subagent should be invoked
tools: tool1, tool2, tool3  # Optional - inherits all tools if omitted
model: sonnet  # Optional - specify model alias or 'inherit'
---

Your subagent's system prompt goes here. This can be multiple paragraphs
and should clearly define the subagent's role, capabilities, and approach
to solving problems.

Include specific instructions, best practices, and any constraints
the subagent should follow.
```

| 字段 | 必需 | 描述 |
| --- | --- | --- |
| `name` | 是 | 使用小写字母和连字符的唯一标识符 |
| `description` | 是 | 子代理目的的自然语言描述 |
| `tools` | 否 | 特定工具的逗号分隔列表。如果省略，继承主线程中的所有工具 |
| `model` | 否 | 用于此子代理的模型。可以是模型别名（sonnet、opus、haiku）或 'inherit' 以使用主对话的模型。如果省略，默认为配置的子代理模型 |

---

#### IV. 使用教程

1. **创建 Sub Agent**: 在 Claude Code 中创建一个 Sub Agent

    {{< figure src="/posts/claude-code-guide/subagent-create-0.png" title="Start to Create a Sub Agent" height="512" width="512" >}}

2. 选择创建的 Sub Agent 属于 `项目级别` / `用户级别`

    {{< figure src="/posts/claude-code-guide/subagent-create-1.png" title="Select Type of the Sub Agent" height="512" width="512" >}}

3. 选择创建方式：`Claude 生成` / `手动配置`

    {{< figure src="/posts/claude-code-guide/subagent-create-2.png" title="Select creation method of Sub Agent" height="512" width="512" >}}

4. 描述 Sub Agent 的功能

    {{< figure src="/posts/claude-code-guide/subagent-create-3.png" title="Desc the function of Sub Agent" height="512" width="512" >}}

5. 选择 Sub Agent 的权限

    {{< figure src="/posts/claude-code-guide/subagent-create-4.png" title="Select the permission of Sub Agent" height="512" width="512" >}}

6. 选择 Sub Agent 的使用的模型

    {{< figure src="/posts/claude-code-guide/subagent-create-5.png" title="Select the model of Sub Agent"  height="512" width="512">}}

7. 设置 Sub Agent 的标识颜色

    {{< figure src="/posts/claude-code-guide/subagent-create-6.png" title="Set the color of Sub Agent"  height="512" width="512">}}

8. 确认、保存创建的 Sub Agent

    {{< figure src="/posts/claude-code-guide/subagent-create-7.png" title="Confirm and save the Sub Agent" height="512" width="512" >}}

---

#### V. Developer's Top Sub Agents

##### **1. Code-Reviewer**

- Name: `code-reviewer`
- Description: Expert code reviewer specializing in code quality, security vulnerabilities, and best practices across multiple languages. Masters static analysis, design patterns, and performance optimization with focus on maintainability and technical debt reduction.
- **Link**: <https://github.com/kydenul/dotfiles/blob/master/claude/agents/code-reviewer.md>

##### **2. Api-Designer**

- Name: `api-designer`
- Description: API architecture expert designing scalable, developer-friendly interfaces. Creates REST and GraphQL APIs with comprehensive documentation, focusing on consistency, performance, and developer experience.
- **Link**: <https://github.com/kydenul/dotfiles/blob/master/claude/agents/api-designer.md>

##### **3. Golang-Pro**

- Name: `golang-pro`
- Description: Expert Go developer specializing in high-performance systems, concurrent programming, and cloud-native microservices. Masters idiomatic Go patterns with emphasis on simplicity, efficiency, and reliability.
- **Link**: <https://github.com/kydenul/dotfiles/blob/master/claude/agents/golang-pro.md>

##### **4. SQL-Pro**

- Name: `sql-pro`
- Description: Expert SQL developer specializing in complex query optimization, database design, and performance tuning across PostgreSQL, MySQL, SQL Server, and Oracle. Masters advanced SQL features, indexing strategies, and data warehousing patterns.
- **Link**: <https://github.com/kydenul/dotfiles/blob/master/claude/agents/sql-pro.md>

##### **5. MicroServices-Architect**

- Name: microservices-architect
- Description: Distributed systems architect designing scalable microservice ecosystems. Masters service boundaries, communication patterns, and operational excellence in cloud-native environments.
- **Link**: <https://github.com/kydenul/dotfiles/blob/master/claude/agents/microservices-architect.md>

> [!TIP]
> 更多的 Sub Agent 请查看[agents](https://github.com/kydenul/dotfiles/blob/master/claude/agents)

### References

- [Claude Code Docs](https://code.claude.com/docs/zh-CN/overview)
- [Slash Commands](https://code.claude.com/docs/en/slash-commands)
- [⚡️10 Claude Code Subagents Every Developer Needs in 2025](https://dev.to/necatiozmen/10-claude-code-subagents-every-developer-needs-in-2025-2ho)


---

> Author: [kyden](https://github.com/kydenul)  
> URL: http://localhost:1313/posts/ea6aa43/  

