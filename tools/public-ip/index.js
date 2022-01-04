"use strict";

(async () => {
  const publicIp = await import("public-ip");
  console.log(await publicIp.default.v4());
})();
