local baseURL = 'http://www.myurl.com'
-- This is gonna be the website URL
local GETKey = ''
-- Will attempt to match with get key on your php server (example: myurl.php?key=GETKey)
local POSTKey = ''
-- Will attempt to match with post key on your php server
-- attempt to validate the POSTKey

-- Action and parameters are based on what you put in the receiver php server
local main = function(action,parameters)
	local ar = {     -- make sure to actually put the parameters in an array even if there is only one of them
		Validate = POSTKey,
		Action = action,
		Parameter1 = parameters[1],
		Parameter2 = parameters[2],
	}
	return game:GetService'HttpService':PostAsync(string.format('%s/transfer.php?key=%s',baseURL,GETKey),game:GetService'HttpService':JSONEncode(ar))
end
