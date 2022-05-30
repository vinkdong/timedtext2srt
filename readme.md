# timedtext2srt
Youtube timedtext to srt format
将Youtube timedtext json字幕转换为SRT格式字幕

## Install

```bash
go install
```

## Usage

- Openfile directly
```bash
timedtext2srt timedtext.json
```

- Openfile with flags and save srt to file
```bash
timedtext2srt -i timedtext.json -o timedtext.srt
```

- Open timedtext and save srt file via pipe
```bash
cat timedtext.json | timedtext2srt > timedtext.srt
```

## 安装方式
- 源码编译安装
```bash
go install
```

- 下载附件安装
[Mac M1芯片](https://github.com/vinkdong/timedtext2srt/releases/download/0.1.0/timedtext2srt-darwin-m1)
[Mac x86_64位](https://github.com/vinkdong/timedtext2srt/releases/download/0.1.0/timedtext2srt-darwin-x86_64)
[Windows](https://github.com/vinkdong/timedtext2srt/releases/download/0.1.0/timedtext2srt.exe)
[Linux](https://github.com/vinkdong/timedtext2srt/releases/download/0.1.0/timedtext2srt-linux)

## 使用方式

- 直接打开文件
```bash
timedtext2srt timedtext.json
```

- 使用参数打开文件和存储srt文件
```bash
timedtext2srt -i timedtext.json -o timedtext.srt
```

- 使用pipe打开json文件和保存srt文件
```bash
cat timedtext.json | timedtext2srt > timedtext.srt
```