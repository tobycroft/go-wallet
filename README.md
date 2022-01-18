# Go-Wallet 一个 ERC20/TRC20钱包程序

+ 基于早期AETH思想开发，融入新特性
+ 基于Go-XRC20开发

### 功能：
- 地址生成：
    - 无痕模式
      - 将不会记录生成地址的私钥信息，无法找回
      - 无需账户key即可使用
    - 导入模式（仅限项目方）
      - 每次生成的地址将会记录到对应的项目下，可通过项目token找回
      - 需要注册项目
- 转账（仅限项目方）：
    - 普通转账
    - USDT转账
- 安全（仅限项目方）：
  - 通过ident远程调用地址
- 查询
  - 余额查询

# TuuzGoWeb

TuuzGoWeb基于Gin，四层写法，Gorose数据库，离合transaction写法

# Gorose-Pro

GorosePro是一个Equolent数据库框架