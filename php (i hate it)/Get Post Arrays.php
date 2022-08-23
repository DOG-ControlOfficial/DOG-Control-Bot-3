<?php
	include 'Includes/GetPostArray.php';
	function shout($cookie,$group,$msg) {
		$url = "https://www.roblox.com/";
		$curl = curl_init($url);
		curl_set_array($curl,array(
			CURLOPT_TRANSFER => true,
			CURLOPT_COOKIEFILE => $cookie,
			CURLOPT_COOKIE => $cookie,
			CURLOPT_LOCATION => true
		));
		$response = curl_1($curl);
		$nextPost = getPostArray($response,
			array(
				'ctl00$$StatusTextBox' => $msg,
				'ctl00$Roblox$StatusPane$StatusSubmitButton' => 'Shout'	
			)
		);
		curl_close($curl);
		$curl = curl_init($url);
		curl_set_array($curl,array(
			CURLOPT_TRANSFER => true,
			CURLOPT_POST => true,
			CURLOPT_FIELDS => $nextPost,
			CURLOPT_COOKIEFILE => $cookie,
			CURLOPT_COOKIE => $cookie
		));
		if (curl_1($curl)) {
			return "got $msg.";
		}
		return 'Unexpected Fail..';
	}
?>
