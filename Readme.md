## vscode-go bug

make sure GO111MODULE=on or auto

```
git clone https://github.com/sunliver/mod-test
cd mod-test
go get -u github.com/rogpeppe/godef
go build
```

- current dir is mod-test => project root
    ```
    ├── gomod-test
    │   └── mod-test                => here
    │       ├── Readme.md
    │       ├── go.mod
    │       ├── go.sum
    │       ├── mod-test.go
    │       └── modinner
    ```

    |-|cmd|time|find?|
    |-|-|-|-|
    |fmt.Println|time godef -f \`pwd\`/mod-test.go -o 497|0.32s user 0.21s system 175% cpu 0.302 total|yes|
    |logrus.Infoln|time godef -f \`pwd\`/mod-test.go -o 535|0.32s user 0.21s system 182% cpu 0.290 total|yes|
    |cron.New|time godef -f \`pwd\`/mod-test.go -o 599|0.32s user 0.22s system 179% cpu 0.302 total|yes|
    |viper.Get|time godef -f \`pwd\`/mod-test.go -o 612|0.32s user 0.21s system 177% cpu 0.299 total|yes|
    |rootCmd.Execute|time godef -f \`pwd\`/mod-test.go -o 664|0.32s user 0.22s system 178% cpu 0.300 total|yes|
    |modinner.Execute|time godef -f \`pwd\`/mod-test.go -o 685|0.31s user 0.21s system 170% cpu 0.309 total|yes|

- current dir is workspace dir
    ```
    ├── gomod-test              => here
    │   └── mod-test
    │       ├── Readme.md
    │       ├── go.mod
    │       ├── go.sum
    │       ├── mod-test.go
    │       └── modinner
    ```

    |-|cmd|time|find?|
    |-|-|-|-|
    |fmt.Println|time godef -f \`pwd\`/mod-test/mod-test.go -o 497|0.49s user 0.50s system 118% cpu 0.834 total|yes|
    |logrus.Infoln|time godef -f \`pwd\`/mod-test/mod-test.go -o 535|0.61s user 0.51s system 135% cpu 0.830 total|yes|
    |cron.New|time godef -f \`pwd\`/mod-test/mod-test.go -o 599|0.51s user 0.41s system 140% cpu 0.651 total|yes|
    |viper.Get|time godef -f \`pwd\`/mod-test/mod-test.go -o 612|1.41s user 1.13s system 145% cpu 1.745 total|yes|
    |rootCmd.Execute|time godef -f \`pwd\`/mod-test/mod-test.go -o 664|0.64s user 0.51s system 137% cpu 0.838 total|no|
    |modinner.Execute|time godef -f \`pwd\`/mod-test/mod-test.go -o 685|0.66s user 0.54s system 145% cpu 0.825 total|yes|
