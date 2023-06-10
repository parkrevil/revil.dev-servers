import { Settings as LuxonSettings } from 'luxon';

import { bootstrap } from './bootstrap';

LuxonSettings.defaultZone = 'Asia/Seoul';

async function start() {
  await bootstrap();
}

start();
