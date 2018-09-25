## Cargo Chaincode

### chaincode method

#### ● createContract - 화물예약의뢰등록 (argc = 9)
<pre><code>Args :
Args[0]- 날짜(ex:20180909),
Args[1]- 화물톤수,
Args[2]- 거리(km), 
Args[3]- 화물의뢰비용(원),
Args[4]- 날짜(yyyy-mm-dd),
Args[5]- 화물등록사용자아이디, 
Args[6]- 운전자아이디,
Args[7]- 수취자아이디,
Args[8]- 상태(SUCCESS,FAIL,YET,COMPLETE)
</code></pre>

#### ● queryCargo - 화물예약의뢰리스트 출력(범위) (argc = 1)
<pre><code>Args :
Args[0]- 날짜(ex:20180909)
ex) 20180909이라면 2018-06-06 ~ 2018-09-09까지의 모든 의뢰리스트 출력,
    20180926이라면 2018-06-06 ~ 2018-09-09까지의 모든 의뢰리스트 출력
</code></pre>


#### ● cancelContract - 화물예약의뢰취소 (argc = 1)
<pre><code>Args :
Args[0]- 해당되는 키값(ex:20180909_1)
</code></pre>

#### ● signContract - 화물의뢰계약 (argc = 2)
<pre><code>Args :
Args[0]- 해당되는 키값(ex:20180909_1),
Args[1]- 운전자아이디
</code></pre>

#### ● completeContract - 수령확인 (argc = 1)
<pre><code>Args :
Args[0]- 해당되는 키값(ex:20180909_1)
</code></pre>

#### ● queryMylist - 특정아이디와 관련한 화물의뢰리스트출력 (argc = 1)
<pre><code>Args :
Args[0]- 사용자아이디
</code></pre>

#### ● queryPoint - 특정계정의 포인트를 조회(argc = 1)
<pre><code>Args :
Args[0]- 사용자아이디
</code></pre>

#### ● addPoint - 특정계정에 포인트를 추가(argc = 2)
<pre><code>Args :
Args[0]- 사용자아이디,
Args[1]- 금액(원)
</code></pre>

#### ● createUser - 계정생성(argc = 2)
<pre><code>Args :
Args[0]- 사용자아이디,
Args[1]- 유저이름
</code></pre>

#### ● subtractPoint - 특정계정에서 포인트 차감(argc = 2)
<pre><code>Args :
Args[0]- 사용자아이디,
Args[1]- 금액
</code></pre>
