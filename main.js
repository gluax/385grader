const axios = require('axios');
const fs = require('fs');
const path = require('path');
const rimraf = require('rimraf');
const unzip = require('unzip');
require('dotenv').config();

async function fetchAss() {
  const { data } = await axios.get(`https://${process.env.CANVAS_API_DOMAIN}/api/v1/courses/${process.env.COURSE_ID}/assignments/${process.env.ASSIGNMENT_ID}/submissions?zip=1&access_token=${process.env.CANVAS_API_KEY}&per_page=1000`);
  const test = [data[0], data[1], data[2]];
  test.map(async ({ attachments, user_id }) => {
    if (attachments) {
      const attachment = attachments.reduce((acc, cur) => {
        if (!acc) {
          return cur;
        }
        return new Date(acc.created_at) > new Date(cur.created_at) ? acc : cur;
      });
      try {
        const url = `${attachment.url}`;
        console.log(url);
        const fn = String(user_id) + '.zip';
        const file = fs.createWriteStream(fn);
        const r = await axios({
          method: 'GET',
          url: url,
          responseType: 'stream'
        });
        r.data.pipe(file);
        r.data.on('finish', () => file.close());
        r.data.on('error', (err) => console.log(err));
        console.log(path.resolve(__dirname, fn));
        const zip = path.resolve(__dirname, fn);
        const folder = path.resolve(__dirname, 'ass');
        //fs.mkdirSync(folder);
        fs.unlinkSync(zip, (err) => console.log(err));
        rimraf(folder, () => {console.log('deleted dir');});
      } catch(err) {
        console.log(err);
      }
      
      return;
    }
  });
}

fetchAss().catch((err)=>console.log(err));
console.log('finish');
