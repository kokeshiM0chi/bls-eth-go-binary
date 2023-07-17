```shell
git pull upstream v1.31.0
```

```shell
% git pull upstream v1.31.0
remote: Enumerating objects: 44, done.
remote: Counting objects: 100% (44/44), done.
remote: Compressing objects: 100% (17/17), done.
remote: Total 44 (delta 15), reused 44 (delta 15), pack-reused 0
Unpacking objects: 100% (44/44), 8.80 MiB | 183.00 KiB/s, done.
From github.com:herumi/bls-eth-go-binary
 * tag               v1.31.0    -> FETCH_HEAD
Updating f61456b..ecb2d6d
Fast-forward
 bls/lib/android/arm64-v8a/libbls384_256.a   | Bin 0 -> 9482186 bytes
 bls/lib/android/armeabi-v7a/libbls384_256.a | Bin 0 -> 5153136 bytes
 bls/lib/android/x86_64/libbls384_256.a      | Bin 0 -> 9100656 bytes
 bls/lib/darwin/amd64/libbls384_256.a        | Bin 0 -> 8364112 bytes
 bls/lib/darwin/arm64/libbls384_256.a        | Bin 0 -> 625240 bytes
 bls/lib/ios/libbls384_256.a                 | Bin 0 -> 15045200 bytes
 bls/lib/iossimulator/libbls384_256.a        | Bin 0 -> 13207720 bytes
 bls/lib/linux/amd64/libbls384_256.a         | Bin 0 -> 916592 bytes
 bls/lib/linux/arm64/libbls384_256.a         | Bin 0 -> 847942 bytes
 bls/lib/windows/amd64/libbls384_256.a       | Bin 0 -> 1351986 bytes
 10 files changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 bls/lib/android/arm64-v8a/libbls384_256.a
 create mode 100644 bls/lib/android/armeabi-v7a/libbls384_256.a
 create mode 100644 bls/lib/android/x86_64/libbls384_256.a
 create mode 100644 bls/lib/darwin/amd64/libbls384_256.a
 create mode 100644 bls/lib/darwin/arm64/libbls384_256.a
 create mode 100644 bls/lib/ios/libbls384_256.a
 create mode 100644 bls/lib/iossimulator/libbls384_256.a
 create mode 100644 bls/lib/linux/amd64/libbls384_256.a
 create mode 100644 bls/lib/linux/arm64/libbls384_256.a
 create mode 100644 bls/lib/windows/amd64/libbls384_256.a

```