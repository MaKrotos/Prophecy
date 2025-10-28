const qrRoutes = [
  {
    path: "/scan",
    name: "scan",
    component: () => import("../views/ScanView.vue"),
    meta: {
      title: "Сканер",
      order: 5,
    },
  },
  {
    path: "/session/my-qr/:id",
    name: "session-my-qr",
    component: () => import("../views/SessionMyQRView.vue"),
    meta: {
      title: "Мой QR код",
      order: 14,
    },
  },
  {
    path: "/session/qr-scanner",
    name: "session-qr-scanner",
    component: () => import("../views/SessionQRScannerView.vue"),
    meta: {
      title: "Сканер QR кодов",
      order: 15,
    },
  },
];

export default qrRoutes;