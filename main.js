<<<<<<< HEAD
const fs = require('fs');
const path = require('path');

const axios = require('axios');
const cmd = require('node-cmd');
=======
const axios = require('axios');
const fs = require('fs');
const path = require('path');
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
const rimraf = require('rimraf');
const unzip = require('unzip');
require('dotenv').config();

async function fetchAss() {
  const { data } = await axios.get(`https://${process.env.CANVAS_API_DOMAIN}/api/v1/courses/${process.env.COURSE_ID}/assignments/${process.env.ASSIGNMENT_ID}/submissions?zip=1&access_token=${process.env.CANVAS_API_KEY}&per_page=1000`);
<<<<<<< HEAD
  const test = [data[0], data[1], data[2], data[3], data[4], data[5]];
=======
  const test = [data[0], data[1], data[2]];
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
  test.map(async ({ attachments, user_id }) => {
    if (attachments) {
      const attachment = attachments.reduce((acc, cur) => {
        if (!acc) {
          return cur;
        }
        return new Date(acc.created_at) > new Date(cur.created_at) ? acc : cur;
      });
<<<<<<< HEAD

      try {
        const url = `${attachment.url}`;
=======
      try {
        const url = `${attachment.url}`;
        console.log(url);
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
        const fn = String(user_id) + '.zip';
        const file = fs.createWriteStream(fn);
        const r = await axios({
          method: 'GET',
          url: url,
          responseType: 'stream'
        });
<<<<<<< HEAD
        
        r.data.pipe(file);
        r.data.on('finish', () => file.close());
        r.data.on('error', (err) => console.log(err));
        
        const zip = path.resolve(__dirname, fn);
        const folder = path.resolve(__dirname, String(user_id));

        fs.createReadStream(zip).pipe(unzip.Extract({ path: folder}).on('close', () => {
          cmd.get(
            `cp ${process.env.TEST_SCRIPT} ${folder}/`,
            (err, data, stderr) => {
              console.log(err);
            }
          );
          
          const isDirectory = source => fs.lstatSync(source).isDirectory();
          const getDirectories = source =>
                fs.readdirSync(source).map(name => path.join(source, name)).filter(isDirectory);
          const subdirs = getDirectories(folder);
          if(subdirs.length) {
            cmd.get(
              `mv ${subdirs[0]}/* ${folder}`,
              (err, data, stderr) => {
                rimraf(subdirs[0], () => {console.log('del');});
                
              }
            );
          }

          cmd.get(`cp ${process.env.TEST_SCRIPT} ${folder}`);

          cmd.get(`cd ${folder}
bash ${process.env.TEST_SCRIPT}
cd ..`,
                  (err, data, stderr) => {
                    const res = data.split('\n');
                    if(data[0].indexOf('done') === -1) {
                      console.log('compilation failed 0');
                      return;
                    }

                    let td = `Compiling gcd.cpp...done

Running test 1...success
Running test 2...success
Running test 3...success
Running test 4...success
Running test 5...success
Running test 6...success
Running test 7...success
Running test 8...success
Running test 9...success
Running test 10...success
Running test 11...failure

Expected________________________________________________________________________
Iterative: gcd(48, -20) = 4
Recursive: gcd(48, -20) = 4
Received________________________________________________________________________
Iterative: gcd(48, -20) = 4

Running test 12...failure

Expected________________________________________________________________________
Iterative: gcd(-8, 80) = 8
Recursive: gcd(-8, 80) = 8
Received________________________________________________________________________
Iterative: gcd(-8, 80) = 8


Total tests run: 12
Number correct : 10
Percent correct: 83.33

Cleaning project...done`;

                    td.split('\n').map((line, index) => {
                      if(line.indexOf(failure) > -1) {
                        
                      }
                    });

                    
                  });
          fs.unlinkSync(zip, (err) => console.log(err));
          
        }));

        
        
=======
        r.data.pipe(file);
        r.data.on('finish', () => file.close());
        r.data.on('error', (err) => console.log(err));
        console.log(path.resolve(__dirname, fn));
        const zip = path.resolve(__dirname, fn);
        const folder = path.resolve(__dirname, 'ass');
        //fs.mkdirSync(folder);
        fs.unlinkSync(zip, (err) => console.log(err));
        rimraf(folder, () => {console.log('deleted dir');});
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
      } catch(err) {
        console.log(err);
      }
      
<<<<<<< HEAD
      
=======
      return;
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
    }
  });
}

<<<<<<< HEAD
fetchAss();
=======
fetchAss().catch((err)=>console.log(err));
console.log('finish');
>>>>>>> a2e735d6e1b103c9ce97d2595aa32951cd806a9d
