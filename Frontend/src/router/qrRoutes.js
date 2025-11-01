const qrRoutes = [
  {
    path: "/scan",
    name: "scan",
    component: () => import("../views/ScanView.vue"),
    meta: {
      title: "Сканер",
    },
  },
  {
    path: "/session/my-qr/:id",
    name: "session-my-qr",
    component: () => import("../views/SessionMyQRView.vue"),
    meta: {
      title: "Мой QR код",
    },
  },
  {
    path: "/session/qr-scanner",
    name: "session-qr-scanner",
    component: () => import("../views/SessionQRScannerView.vue"),
    meta: {
      title: "Сканер QR кодов",
    },
  },
];

export default qrRoutes;