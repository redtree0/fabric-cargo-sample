## Hyperledger Fabric Cargo Sample - Dev_ljs

### version
<pre><code>node.js - 8.11.1
docker-machine - 0.14.0
docker - v18.03.0-ce
mysql 8.0
</code></pre>

### 회원가입,로그인,로그아웃 구현 (CA와 연동)
<pre><code>1) 회원가입시 [mysql] blockchain(database) => users에 회원정보저장
2) 회원정보가 저장된후 fabric_CA에 register
3) 로그인시 [mysql] users table과 먼저 비교
4) 3번진행후 enroll을 통해 인증서와 개인키를 발급받고 상태저장소에 저장
</code></pre>

### 사용법
<pre><code>1) mysql을 먼저설치한후 (유의사항 : Authentication Method 설정시 legacy로 설정하세요)
2) mysql root 패스워드는 konyang으로 해주세요.
3) test.sql파일을 적용한다.(mysql -uroot -pkonyang < 경로\test.sql)
4) cargo-app폴더가 들어가서 npm install 수행
5) ./startFabric.sh
6) node registerAdmin.js
7) node ./bin/www
8) http://localhost:8000/login 접속
</code></pre>

[![Alt text for your video](https://i.ytimg.com/vi/g0MRWP9w86Q/hqdefault.jpg)](https://www.youtube.com/watch?v=g0MRWP9w86Q)




