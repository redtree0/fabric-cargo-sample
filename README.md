## Hyperledger Fabric Cargo Sample

[facar 예제](http://hyperledger-fabric.readthedocs.io/en/release-1.1/write_first_app.html)를 기반으로 만든 예제입니다.

### version
<pre><code>node.js - 6.11.4
docker-machine - 0.14.0
docker - v18.05.0-ce 
</code></pre>

### Build and Test

####  ["먼저 시작해 보는 블록체인 03 – Hyperledger Fabric 개발 환경 구성"](https://developer.ibm.com/kr/developer-%EA%B8%B0%EC%88%A0-%ED%8F%AC%EB%9F%BC/2017/01/26/blockchain-basic-03-build_development_environment/) 참고
<pre><code>docker-machine create --driver virtualbox blockchain </code></pre>
<pre><code>docker-machine env blockchain </code></pre>
<pre><code>eval $(docker-machine env blockchain) </code></pre>


#### cargo-app 내에 웹 어플리케이션 구현
<pre><code>cd cargo-app</code></pre>

<pre><code>npm install</code></pre>


#### Docker-Compose 실행
<pre><code>./startFabric.sh</code></pre>

 
#### Hyperledger Fabric CA 계정 등록
<pre><code>node registerAdmin.js</code></pre>

<pre><code>node registerUser.js</code></pre>


#### 웹 어플리케이션 실행
<pre><code>node ./bin/www</code></pre>

