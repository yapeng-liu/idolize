# jaeger

[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](grafana-promtail-loki)

- [组件说明](#组件说明)

---

### 组件说明
   *  loki 负责存储日志和处理查询
   *  promtail 负责收集日志并发送日志至loki
   *  grafana 进行UI展示
### docker-compose
~~~
  grafana:
    image: grafana/grafana:latest
    ports:
      - "3000:3000"
    networks:
      - docoi-test
  promtail:
    image: grafana/promtail:latest
    volumes:
      - /data/app/docoi_test/logs:/var/log/
    command: -config.file=/etc/promtail/config.yml
    networks:
      - docoi-test
  loki:
    image: grafana/loki:latest
    ports:
      - "3100:3100"
    command: -config.file=/etc/loki/local-config.yaml
    networks:
      - docoi-test
    logging:
      driver: "json-file"
      options:
        max-size: "5g"
~~~
### 登录
    * http://ip:3000

