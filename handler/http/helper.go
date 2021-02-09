package http

import "net/http"

func setHeader(req *http.Request) {
	req.Header.Set("Cookie", "xq_a_token=176b14b3953a7c8a2ae4e4fae4c848decc03a883; xqat=176b14b3953a7c8a2ae4e4fae4c848decc03a883; xq_r_token=2c9b0faa98159f39fa3f96606a9498edb9ddac60; xq_id_token=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOi0xLCJpc3MiOiJ1YyIsImV4cCI6MTYxMzQ0MzE3MSwiY3RtIjoxNjEyNjIwMTA1ODk4LCJjaWQiOiJkOWQwbjRBWnVwIn0.kpeZvJAZJp7azwXMwdRch45UiiTSuUDD91q31PXDcTYils9OoVNszM2NmV_E9NWJw0tqM35dRctbbmo3mznfp38qdHEDK0OK1fJJ_rkbS8UYnlUBKRJjkzFdHA2D6x02IbxGbm_AEN70RrauLHpin3bb5gYGNFppP-YYUmSgLwAaVLOXPqHWANp0qGQ7YEGKAk6r65HUCtpyqod--1eoiETB-db7BldkqieMozmLGacsS7UyarjrSC0tglYOdDsew9bxKyrozpxf955wFkBm1233tXQ8CeqCbR2xjZk9VpnlIg1yyjgFq1wNXZCnFCTh5B17NbAWJ4X49N1wp87_Dg; u=691612620134423; device_id=24700f9f1986800ab4fcc880530dd0ed; Hm_lvt_1db88642e346389874251b5a1eded6e3=1612620136,1612620150,1612689387,1612702614; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1612702614")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://xueqiu.com")
}
