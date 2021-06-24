# mag (Mac Address Generator)

[![Test](https://github.com/rnazmo/mag/actions/workflows/test.yml/badge.svg)](https://github.com/rnazmo/mag/actions/workflows/test.yml)

Generate a random MAC address (EUI-48).

## TODO

- [ ] Setup lint
  - [ ] Run lint on local
  - [ ] Run lint on CI
- [ ] Setup unit-test
  - [ ] Run unit-test on local
  - [ ] Run unit-test on CI
- [ ] Add tests

- [ ] Add support for specifying the address range
  - Allows you to specify a generated address range.
  - default: 00:00:00:00:00:00 ~ ff:ff:ff:ff:ff:ff
- [ ] Add support for cli option
- [ ] Web 版作る (ブラウザからも使いたい)
- [ ] oui リストの参照元を変える
  - before: http://standards-oui.ieee.org/oui/oui.csv
    - IEEE の公開している MAC Address Block Large (MA-L)
    - 29247 行 (20210130 時点)
    - 単なるCSVなのでパースが楽
    - HTTPリクエスト時の応答が遅い
  - after: https://gitlab.com/wireshark/wireshark/raw/master/manuf
    - Wireshark の公開しているリスト
    -  41751 行 (20210130 時点)
    - 情報量が多いがパースが面倒そう
