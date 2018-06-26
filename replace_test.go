package main

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	content := `.logo {
		height: 8rem;
		background: url(http://www.youhopelife.com/imgs/logo.png) no-repeat;
	}
	
	.sec {
		margin-top: 1.8rem;
		position: relative;
	}
	
	.search-content {
		position: relative;
	}
	
	.search {
		width: 80%;
		height: 3rem;
		padding-left: 1rem;
		border: 2px solid #008CE1;
	}
	
	.shop,
	.goods {
		display: block;
		width: 5rem;
		height: 2.5rem;
		font-size: 1.4rem;
		line-height: 2.5rem;
		color: #fff;
		text-align: center;
		float: left;
	}
	
	.shop {
		background: #008CE1;
	}
	
	.goods {
		color: #666;
	}
	
	.icon-search {
		background: #008CE1;
		font-size: 2rem;
		position: absolute;
		top: 32.5%;
		margin-left: -2.5rem;
		color: #ffffff;
		height: 3rem;
		width: 4.3rem;
		text-align: center;
		line-height: 3rem;
	}
	
	.hot-search {
		margin-top: 0.5rem;
	}
	
	.hot-text {
		font-weight: bold;
	}
	
	.hot-search a {
		color: #333;
	}
	
	.icon-people {
		width: 2.4rem;
		height: 2.8rem;
		display: inline-block;
		vertical-align: middle;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
		background-repeat: no-repeat;
		background-position: -1px -15px;
	}
	
	.icon-down {
		width: 1.8rem;
		height: 1.4rem;
		display: inline-block;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
		background-repeat: no-repeat;
		background-position: 6px -110px;
	}
	
	.my-yh,
	.my-car {
		height: 3.4rem;
		width: 14rem;
		line-height: 3.4rem;
		background: #f9f9f9;
		border: 1px solid #d3d3d3;
		color: #676767;
		font-size: 1.4rem;
		margin-top: 2.2rem;
	}
	
	.my-car {
		margin-left: 1rem;
		width: 12.5rem;
	}
	
	.icon-car {
		width: 2.4rem;
		height: 2.8rem;
		display: inline-block;
		vertical-align: middle;
		background-position: 1px -75px;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
		background-repeat: no-repeat;
	}
	
	.head-nav {
		height: 5.4rem;
		border-bottom: 2px solid #018DE2;
		margin-top: 1rem;
		position: relative;
	}
	
	.many-lists {
		position: absolute;
		width: 36.5rem;
		height: 46.8rem;
		left: 19.88%;
		top: 5.3rem;
		background: rgba(0,0,0,.4);
		z-index: 1000;
		display: none;
	}
	
	.ul-item {
		display: none;
	}
	
	.many-lists li {
		float: left;
		color: #fff;
		font-size: 1.4rem;
		border-right: 1px solid #d3d3d3;
		margin-right: 1rem;
		margin-top: 1.5rem;
		padding-right: 1rem;
	}
	
	.many-lists li:hover{
		color: rgb(211,126,60);
		cursor: pointer;
	}
	
	.nav-list {
		position: absolute;
		width: 17.52rem;
		background: #404040;
		left: 8.37%;
		top: 5.3rem;
		overflow: hidden;
		z-index: 1000;
		display: none;
	}
	
	.nav-list a {
		display: block;
		color: #fff;
		font-size: 1.4rem;
		text-indent: 4rem;
		height: 5.2rem;
		line-height: 5.2rem;
	}
	
	.nav-list li {
		width: 17.52rem;
		height: 5.2rem;
		border-bottom: 1px solid #6a6a6a;
		position: relative;
	}
	
	.icon-i {
		display: inline-block;
		position: absolute;
		right: 0;
		top: 0;
		width: 5.2rem;
		height: 5.2rem;
		background-position: 25px -444px;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
		background-repeat: no-repeat;
	}
	
	.type-content {
		width: 17.5rem;
		height: 5.4rem;
		background: #018DE2;
		font-size: 1.6rem;
		text-align: center;
		line-height: 5.4rem;
		cursor: pointer;
	}
	
	.types {
		color: #fff;
	}
	
	.types:hover {
		color: #fff;
	}
	
	.icon-xuanzhuan {
		width: 3rem;
		height: 3rem;
		display: inline-block;
		vertical-align: middle;
		background-position: 6px -176px;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
		background-repeat: no-repeat;
		transition: all 0.3s ease;
	}
	
	.icon-xuanzhuan:hover {
		transform: rotate(-180deg);
	}
	
	.more-type li {
		font-size: 1.6rem;
		color: #676767;
		float: left;
		height: 4.4rem;
		line-height: 5.4rem;
		text-align: center;
		padding: 0 4rem;
		cursor: pointer;
	}
	
	.online {
		font-size: 1.6rem;
		color: #676767;
	}
	
	.online-content {
		height: 4.4rem;
		line-height: 5.4rem;
		text-align: center;
	}
	
	.join {
		margin-top: 1rem;
		margin-bottom: 3rem;
		margin-left: 9%;
		width: 80%;
		text-align: center;
	}
	
	.join img {
		margin-right: 1rem;
	}
	
	.marks {
		margin-left: 8.7%;
		width: 80%;
		border-bottom: 3px solid #018DE2;
	}
	
	.choice {
		height: 3.2rem;
		margin-bottom: 2rem;
		margin-left: -1.5rem;
	}
	
	.text1,
	.text2 {
		font-size: 3rem;
		line-height: 3rem;
	}
	
	.text1 {
		color: #FE6E00;
	}
	
	.text2 {
		color: #008CE1;
	}
	
	.text3 {
		color: #666;
		font-size: 1.6rem;
		line-height: 4rem;
	}
	
	.mLeft {
		margin-left: 8.7%;
		width: 80%;
		position: relative;
	}
	
	.come {
		background: #fff;
		border: 1px solid #D4D4D4;
		border-top: 0;
		padding: 10px;
		margin-bottom: 1rem;
	}
	
	.ad {
		margin-bottom: 3rem;
	}
	
	.ad img {
		width: 100%;
	}
	
	.eat {
		border-bottom: 3px solid #018DE2;
		display: flex;
		flex-direction: row;
	}
	
	.titles {
		font-weight: normal;
		width: 25%;
		height: 3.12rem;
		line-height: 3.12rem;
		font-size: 2.6rem;
		color: #008CE1;
		margin: 0 0 1rem 0;
		position: relative;
	}
	
	.titles span {
		position: absolute;
		left: 4.6rem;
		top: 0.15rem;
	}
	
	.titles img {
		width: 3rem;
		height: 3rem;
	}
	
	.eat li {
		float: left;
		padding-right: 2.5rem;
		cursor: pointer;
	}
	
	.eatUl {
		line-height: 3.12rem;
		font-size: 1.6rem;
		color: #666;
		position: relative;
	}
	
	.more {
		line-height: 3.12rem;
		font-size: 1.6rem;
		color: #666;
		position: absolute;
		right: 3rem;
	}
	
	.lastLi {
		position: absolute;
		width: 5%;
		height: 0.3rem;
		background: #FF6D00;
		left: 0;
		bottom: -1.3rem;
	}
	
	.lastLi:after {
		position: absolute;
		top: -1.3rem;
		left: 50%;
		content: "";
		width: 0;
		height: 0;
		border-width: 0.7rem;
		border-style: solid;
		border-color: transparent;
		border-bottom-color: #FF6D00;
		margin-left: -0.6rem;
	}
	
	.foot {
		width: 100%;
	}
	
	.foot li {
		float: left;
		width: 20%;
		text-align: center;
		font-size: 2.6rem;
		color: #666;
		line-height: 5rem;
	}
	
	.ele1 {
		display: inline-block;
		background-position: -116px -250px;
		width: 4.5rem;
		height: 4.5rem;
		vertical-align: middle;
		margin: 0 2rem 0 0;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
	}
	
	.ele2 {
		display: inline-block;
		background-position: -114px 326px;
		width: 4.5rem;
		height: 4.5rem;
		vertical-align: middle;
		margin: 0 2rem 0 0;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
	}
	
	.ele3 {
		display: inline-block;
		background-position: -113px 166px;
		width: 4.5rem;
		height: 4.5rem;
		vertical-align: middle;
		margin: 0 2rem 0 0;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
	}
	
	.ele4 {
		display: inline-block;
		background-position: -114px 408px;
		width: 4.5rem;
		height: 4.5rem;
		vertical-align: middle;
		margin: 0 2rem 0 0;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
	}
	
	.ele5 {
		display: inline-block;
		background-position: -116px 475px;
		width: 45px;
		height: 45px;
		vertical-align: middle;
		margin: 0 20px 0 0;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
	}
	
	.top {
		width: 100%;
		padding: 0 1.5rem;
		border-top: 1px solid #e4e4e4;
		border-bottom: 1px solid #e4e4e4;
	}
	
	.top-icons {
		width: 80%;
		min-height: 7rem;
		margin: 1.5rem auto 0 8%;
	}
	
	.center {
		height: 24.6rem;
		position: relative;
		border-bottom: 1px solid #d3d3d3;
	}
	
	.v {
		position: absolute;
		left: 8.7%;
		top: 50%;
		transform: translate(0, -50%);
	}
	
	.ele {
		width: 14.5rem;
		height: 14.5rem;
		margin-right: 3rem;
		margin-top: -3.5rem;
	}
	
	.my-logo {
		margin-right: 10rem;
	}
	
	.my-text {
		font-size: 1.6rem;
		color: #666;
		position: absolute;
		bottom: -1rem;
		left: 32rem;
	}
	
	.my-text span {
		display: inline-block;
		width: 14rem;
		height: 3rem;
		text-align: center;
		float: left;
		line-height: 3rem;
		margin-right: 3.5rem;
	}
	
	.icon-phone {
		display: inline-block;
		vertical-align: -30%;
		width: 3rem;
		height: 3rem;
		background-position: -49px -257px;
		background-image: url(http://www.youhopelife.com/imgs/icons.png);
	}
	
	.foot-nav {
		float: right;
		width: 50rem;
		margin-top: -4rem;
		margin-left: 2rem;
		padding-left: 6rem;
		border-left: 1px solid #d3d3d3;
	}
	
	.foot-nav li {
		width: 25%;
		text-align: left;
	}
	
	.foot-nav h3 {
		margin-bottom: 2.5rem;
	}
	
	.foot-nav h3 a {
		font-size: 1.6rem;
		color: #666;
		font-weight: bold;
	}
	
	.foot-nav h4 {
		margin-bottom: 1.5rem;
	}
	
	.foot-nav h4 a {
		font-size: 1.4rem;
		margin-bottom: 1.5rem;
		color: #666;
	}
	
	.bottom {
		text-align: center;
		margin: 4rem 0 5rem 0;
	}
	
	.copyright {
		font-size: 1.4rem;
		color: #999;
		margin-bottom: 1rem;
	}
	
	.bottom-text {
		font-size: 1.4rem;
		color: #2784C7;
	}
	
	.right-floor {
		width: 4.5rem;
		border: solid 1px #E0E0E0;
		position: fixed;
		right: 7%;
		bottom: 10.5rem;
		overflow: hidden;
		display: none;
		box-sizing: content-box;
	}
	
	.right-floor ul{
		width: 5rem;
		margin: 0;
		padding: 0;
	}
	
	
	.right-floor li {
		width: 4.8rem;
		border-top: solid 1px #E0E0E0;
		padding: 0.5rem 0;
		display: block;
	}
	
	.right-floor li:first-child {
		border: 0;
	}
	
	.right-span {
		color: #018CDF;
	}
	
	.right-floor a {
		display: block;
		font-size: 1.5rem;
		text-align: center;
		color: #666;
	}
	
	.right-floor .right-li:hover {
		background: rgb(0,140,225);
	}
	
	.right-icon-up {
		margin-left: -0.5rem;
	}
	
	.w-slider {
		margin: 0 -1.5rem;
	}
	
	.join-img {
		position: relative;
		height: 36.2rem;
		background: url(http://odho0ke5x.bkt.clouddn.com/goodsimg/1466668206227.jpg) no-repeat scroll center center;;
		background-size: 100% 100%;
	}
	
	.join-img p {
		height: 5rem;
		line-height: 4.5rem;
		font-size: 2.6rem;
		color: #666;
	}
	
	.behind-text {
		color: #ff6c00;;
	}
	
	.jion-v {
		position: absolute;
		top: 50%;
		left: 8.7%;
		margin-top: -10rem;
	}
	
	.select-box {
		width: 82%;
		height: 60rem;
		margin-left: 8.3%;
		position: relative;
	}
	
	.amuser {
		width: 58.9rem;
		height: 40rem;
		border-right: 1px solid #d3d3d3;
		position: absolute;
		top: 50%;
		left: 0;
		margin-top: -20rem;
	}
	
	.amuser2 {
		width: 58.9rem;
		height: 40rem;
		position: absolute;
		top: 50%;
		left: 59rem;
		margin-top: -20rem;
	}
	
	.amuser-cricle {
		float: left;
		width: 40rem;
		height: 40rem;
		border: 1px solid #EC6917;
		border-radius: 50%;
		position: relative;
	}
	
	.amuser-cricle2 {
		float: right;
		width: 40rem;
		height: 40rem;
		border: 1px solid #008CE1;
		border-radius: 50%;
		position: relative;
	}
	
	.amuser-cricle2 .mark5 {
		background: #008CE1;;
	}
	
	.amuser-p-box {
		position: absolute;
		top: 50%;
		left: 50%;
		margin-top: -10.8rem;
		margin-left: -7.2rem;
	}
	
	.mark1 {
		text-align: center;
		height: 5rem;
		font-size: 3.6rem;
		color: #EC6917;
		line-height: 5rem;
	}
	
	.mark2 {
		text-align: center;
		font-size: 2rem;
		color: #666;
		height: 3rem;
		line-height: 3rem;
		margin: 1rem 0;
	}
	
	.mark3 {
		text-align: center;
		font-size: 1.4rem;
		color: #999;
		height: 3rem;
		line-height: 3rem;
	}
	
	.mark4 {
		text-align: center;
		height: 3.5rem;
		margin: 1rem 0 0 0;
	}
	
	.mark4 a {
		height: 3.5rem;
		line-height: 3.5rem;
		display: inline-block;
		background: #EC6917;
		color: #fff;
		font-size: 1.6rem;
		border-radius: 0.5rem;
		padding: 0 3rem;
	}
	
	.logo-box2 {
		position: fixed;
		left: 0;
		top: 0;
		right: 0;
		background: #eeeeee;
		z-index: 2000;
	}
	
	.jf-content {
		width: 85%;
		margin-left: 8%;
		display: none;
	}
	
	.jf-product-list {
		width: 21%;
		margin: 2%;
		display: inline-block;
		border: 1px solid #eee;
	}
	
	.jf-product-list img {
		width: 80%;
		height: 46%;
		margin-left: 10%;
		margin-top: 1rem;
	}
	
	.jf-product-list > p {
		margin-top: 2rem;
		margin-left: 2.5rem;
		color: black;
		font-size: 1.6rem;
	}
	
	.jz span:first-child {
		height: 3rem;
		line-height: 3rem;
		color: #FF6C00;
		font-size: 1.6rem;
		margin-left: 2.5rem;
	}
	
	.jz span:last-child {
		height: 3rem;
		line-height: 3rem;
		font-size: 1.6rem;
		margin-left: 1rem;
		color: #666;
	}
	
	.jf-product-list button {
		color: #fff;
		margin: 2rem 0 2rem 33%;
		border: 0;
		width: 33%;
		height: 3rem;
		line-height: 3rem;
		background: #FF6C00;
		border-radius: 0.5rem;
	}
	
	.tj-content,
	.jx-content {
		width: 81%;
		margin-left: 10%;
		display: none;
	}
	
	.jx-product-list2 {
		width: 28.8rem;
		float: left;
		margin: 2rem;
	}
	
	.sj-name {
		margin-top: 2rem;
	}
	
	.jx-product-list2 img {
		width: 28.7rem;
		height: 15rem;
	}
	
	.jx-product-list2 span {
		margin-left: 1rem;
	}
	
	.tj span {
		margin-left: 2.5rem;
		font-size: 1.6rem;
		color: red;
	}
	
	.jx-i1,
	.jx-i2,
	.jx-i3 {
		display: inline-block;
		background-repeat: no-repeat;
		width: 1.6rem;
		height: 1.7rem;
	}
	
	.jx-i1 {
		background-image: url('http://www.youhopelife.com/imgs/sj.png');
		vertical-align: -12%;
	}
	
	.jx-i2 {
		background-image: url('http://www.youhopelife.com/imgs/icons.png');
		background-repeat: no-repeat;
		vertical-align: -12%;
		background-position: -204px -52px;
	}
	
	.jx-i3 {
		background-image: url('http://www.youhopelife.com/imgs/icons.png');
		background-repeat: no-repeat;
		vertical-align: -12%;
		background-position: -204px -83px;
	}
	
	.tj .yj-p {
		font-size: 1.6rem;
		color: #666;
		text-decoration: line-through;
	}
	
	.tj .xsl {
		font-size: 1.6rem;
		color: #666;
	}
	
	.tj {
		display: flex;
		flex-direction: column;
	}
	
	.jf-product-list .s-name {
		line-height: 3rem;
		height: 3rem;
		text-align: center;
		margin: 1rem 0 0 0;
	}
	
	.f-img {
		width: 31.6rem;
		height: 51.1rem;
	}
	
	.join img {
		width: 22%;
	}
	
	/*.eat-f {*/
		/*height: 48.8rem;*/
		/*overflow-x: scroll;*/
		/*overflow-y: hidden;*/
		/*white-space: nowrap;*/
	/*}*/
	
	@media (max-width: 768px) {
		.city {
			text-align: center;
		}
	
		.more-type li {
			margin-left: -1rem;
			padding: 0 3rem;
		}
	}
	
	@media (max-width: 1250px) {
		.more-type li {
			padding: 0 2rem;
		}
	}
	
	@media (max-width: 1053px) {
		.more-type li {
			padding: 0 1rem;
		}
	}
	
	@media (min-width: 670px) and (max-width: 768px) {
		.search {
			width: 75%;
		}
	
		.icon-search {
			top: 0;
		}
	}
	
	@media (max-width: 1373px) {
		.join img {
			margin-top: 1rem;
		}
	}
	
	@media (max-width: 1100px) {
		.titles {
			width: 15rem;
		}
	
		.eat li {
			padding-right: 2rem;
		}
	}
	
	@media (max-width: 620px) {
		.titles {
			width: 10rem;
		}
	}
	
	@media (min-width: 300px) and (max-width: 480px) {
		.more-type {
			margin: 0 auto;
			width: 100%;
		}
	
		.come img {
			width: 100%;
		}
	}
	
	@media (max-width: 1390px) {
		.center {
			display: none;
		}
	}
	
	@media (max-width: 550px) {
		.join img {
			width: 20%;
			margin-right: 0.5rem;
		}
	}`
	fmt.Println(replaceRem(content, 100))
}
