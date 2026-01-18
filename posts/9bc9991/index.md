# OpenAI Whisper 完全指南：高效实现语音识别与字幕生成


{{< admonition type=abstract title="导语" open=true >}}
在现代内容创作和开发场景中，语音识别和字幕生成已成为提升效率的关键技术。OpenAI Whisper 作为目前最先进的自动语音识别（ASR）系统之一，凭借其卓越的多语言识别能力和高准确度，已成为语音转文字领域的标杆工具。

Whisper.cpp 作为 Whisper 的 C++ 实现，在保持高精度的同时，大幅提升了推理速度，使其在普通硬件上也能流畅运行。无论是视频字幕制作、会议记录转录，还是实时语音识别应用，Whisper 都能提供可靠的解决方案。

本文将带你全面了解 Whisper 和 Whisper.cpp 的各项功能，从基础安装到高级配置，从模型选择到实战应用，帮助你在 macOS 上构建高效的语音识别工作流，实现从音频到字幕的自动化转换。
{{< /admonition >}}

<!--more-->

## OpenAI Whisper 简介

[Whisper](https://openai.com/research/whisper) 是 OpenAI 开源模型，可以**实现识别音频、视频中的人声，并将人声转换为字幕内容**，保存到文件

[whisper-cpp](https://github.com/ggml-org/whisper.cpp) 是基于 Whisper 模型的 C++ 实现，与原版 Python 相比，具有以下优点：

- 🚀 更快速：C++ 实现使推理速度提升 4-5 倍
- 💻 更轻量：不需要 Python、PyTorch 或大型 ML 框架
- 📱 更便携：可在多种装置上运行，包括 CPU 和 GPU 支持

> [!NOTE]
> 主要讲解如何在 MacOS 上使用 Whisper.cpp 模型进行字幕文件的生成

---

### 主要功能

- 即时麦克风语音辨识
- 支持多国语言（含繁体/简体中文）
- 高精度语音转录
- 低延迟和高效能
- 语音活动侦测（VAD）整合
- 时间戳标记（定位词语时间点）
- 输出多种格式（文字、SRT、VTT 字幕）
- 支持 Apple Silicon 原生加速

---

## 安装

#### Homebrew Install

```bash
brew install whisper-cpp ffmpeg sdl2
```

> [!TIP]
> Homebrew 安装适合快速入门，但功能相对有限。
> 如需要最大化效能或启用进阶功能（如 Metal 加速），**建议** 使用源码编译安装

- `ffmpeg`: 依赖库
- `sdl2`: whisper-cpp 的即时转录功能依赖 SDL2 来获取麦克风数据

---

#### 源码编译

1. Whisper

    ```bash
    git clone https://github.com/ggml-org/whisper.cpp.git
    cd whisper.cpp

    # 可选，切换至稳定版，截止本稿写时为 v1.8.3
    git checkout v1.8.3
    ```

2. 启用 Metal 和 Core ML 加速（Apple Silicon 专用）

- 对于 M1/M2/M3 芯片的 Mac，可以启用 Metal 和 Core ML 加速以获得最佳效能：

    ```bash
    # 使用 Metal 加速（Apple Silicon 优化）
    cmake -B build -DWHISPER_METAL=ON -DWHISPER_SDL2=ON

    # 使用 Core ML 加速
    cmake -B build -DWHISPER_COREML=ON -DWHISPER_SDL2=ON

    cmake --build build --config Release

    ```

- 标准编译

    ```bash
    cmake -B build -DWHISPER_SDL2=ON
    cmake --build build --config Release
    ```

> [!TIP]
> 编译完成后，`whisper-stream` / `whisper-cli` 等可执行文件位于 `build/bin` 目录下

---

## 模型选择

| Model Name | Model Size | Mermory Require | Speed | 效果 | Suppurt Chinese | Remark |
| --- | --- | --- | --- | --- | --- | --- |
| tiny | ~75MB | ~390MB | 最快 | 最低 | ✅ 基本支持 | 适合极度受限的装置 |
| base | ~142MB | ~500MB | 快 | 较低 | ✅ 支持 | 适合普通装置 |
| small | ~466MB | ~1.0GB | 中等 | 中等 | ✅ 良好支持 | 平衡速度与准确度 |
| medium | ~1.5GB | ~2.0GB | 较慢 | 较高 | ✅ 良好支持 | 适合专业转录 |
| large-v3 | ~3.1GB | ~5.0GB | 慢 | 最高 | ✅ 最佳支持 | 最高准确度 |
| large-v3-turbo | ~3.0GB | ~4.8GB | 中等 | 较高 | ✅ 最佳支持 | 平衡速度与准确度 |

> [!NOTE]
>
> - 后缀 `-q5_0` 或 `-q8_0` 是量化版本
> - 后缀 `-turbo` 是2024年新增的平衡模型，提供较好的效能和精度平衡

> [**下载链接**](https://huggingface.co/ggerganov/whisper.cpp/tree/main)

> [!TIP]
> 下载相关模型后，将其放入 `models/` 文件夹下即可

---

## Whisper 使用指南

Whisper.cpp 最佳支持 **16kHz、16-bit 单声道 WAV 格式**。如果您的音讯档案是其他格式，需要先进行转换。

### FFmpeg 声音格式转换

```bash
ffmpeg -i ~/Movies/SiliconGrown/voice.mp3 -ar 16000 -ac 1 -c:a pcm_s16le output.wav

ffmpeg -i input.mp3 -ar 16000 -ac 1 -c:a pcm_s16le output.wav
```

- `ffmpeg`: 调用 ffmpeg 程序
- `-i input.mp3`: 指定输入文件为 input.mp3
- `-ar 16000`: 设置音频采样率为 16000 Hz (16 kHz),
  - `ar` 是 audio rate 的缩写
  - `16000` Hz 是语音识别常用的采样率,比 CD 质量的 44100 Hz 低,但足够语音使用
- `-ac 1`: 设置音频声道数为 1(单声道)
  - `-ac`: 是 audio channels 的缩写
  - `1`: 表示单声道(mono), `2` 表示立体声(stereo)
- `-c:a pcm_s16le`: 设置音频编码格式
  - `-c:a`: codec:audio 的缩写,指定音频编解码器
  - `pcm_s16le` 表示:
    - PCM - 脉冲编码调制(未压缩的音频格式)
    - s16 - 16位有符号整数采样
    - le - 小端字节序(Little Endian)
- `output.wav`: 输出文件名

将 MP3 文件转换为 WAV 格式,配置为:16kHz 采样率、单声道、16位 PCM 编码。这种配置常用于语音识别、语音处理等场景,因为它在保证语音质量的同时减小了文件大小。

### 视频格式提取 wav 音频

```bash

ffmpeg -i input.mp4 -ar 16000 -ac 1 -c:a pcm_s16le output.wav
```

- `ffmpeg`: 调用 ffmpeg 程序
- `-i input.mp4`: 指定输入**视频**文件为 input.mp4
- `-ar 16000`: 设置音频采样率为 16000 Hz (16 kHz),
  - `ar` 是 audio rate 的缩写
  - `16000` Hz 是语音识别常用的采样率,比 CD 质量的 44100 Hz 低,但足够语音使用
- `-ac 1`: 设置音频声道数为 1(单声道)
  - `-ac`: 是 audio channels 的缩写
  - `1`: 表示单声道(mono), `2` 表示立体声(stereo)
- `-c:a pcm_s16le`: 设置音频编码格式
  - `-c:a`: codec:audio 的缩写,指定音频编解码器
  - `pcm_s16le` 表示:
    - PCM - 脉冲编码调制(未压缩的音频格式)
    - s16 - 16位有符号整数采样
    - le - 小端字节序(Little Endian)
- `output.wav`: 输出音频文件名

FFmpeg 会自动识别 MP4 中的音频流，然后将其转换为 WAV 格式，保存在 output.wav 文件中。
这种转换通常用于提取视频中的音频,以便进行语音识别、语音处理等操作。

### 使用 `whisper` 命令输出字幕

```bash
./build/bin/whisper-cli \
  -m models/ggml-large-v3.bin \
  -l zh \
  -osrt \
  -of subtitles \
  -f audio.wav
```

- `-m models/ggml-large-v3.bin`: 指定模型文件
- `-l zh`: 设置语言为中文
- `-osrt`: 输出 SRT 格式字幕
- `-of subtitles`: 输出字幕文件
- `-f audio.wav`: 指定音频文件

### 常用指令示例

#### 即使转录

`whisper-mic.sh` 脚本文件：

```bash
#!/bin/bash
./build/bin/whisper-stream \
  -m models/ggml-large-v3-q5_0.bin \
  -l zh \
  -t 4 \
  --step 500 \
  --length 5000 \
  --keep 200 \
  --vad-thold 0.6 \
  --freq-thold 100.0 \
  -ps \
  -kc \
  -f whisper_output.txt
```

#### 离线转录

```bash
./build/bin/whisper-cli \
  -m models/ggml-large-v3-q5_0.bin \
  -f samples/zh.wav \
  -l zh \
  --output-txt \
  --output-srt
```

#### 自动检测语言

```bash
./build/bin/whisper-cli \
  -m models/ggml-large-v3.bin \
  -f samples/audio.wav \
  --auto-language
```

## Reference

- [OpenAI Whisper](https://github.com/openai/whisper)
- [Hugging Face 模型](https://huggingface.co/ggerganov/whisper.cpp/tree/main)
- [在线试用](https://ggml.ai/whisper.cpp/)


---

> Author: [kyden](https://github.com/kydenul)  
> URL: https://kydenul.github.io/posts/9bc9991/  

