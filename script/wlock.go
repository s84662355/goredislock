package script


var wlock string = `

		redis.replicate_commands()
		function string.split(input, delimiter)
			input = tostring(input)
			delimiter = tostring(delimiter)
			if (delimiter=='') then return false end
			local pos,arr = 0, {}

			for st,sp in function() return string.find(input, delimiter, pos, true) end do
			table.insert(arr, string.sub(input, pos, st - 1))
			pos = sp + 1
			end
			table.insert(arr, string.sub(input, pos))
			return arr
		end

		local key   = KEYS[1]
		local expire =  tonumber( ARGV[1] )
		local client_name =  ARGV[2]
		local a = redis.call('TIME')
		local cur_timestamp =  tonumber( a[1]  )
		local result=0
		local lockdata =  expire + cur_timestamp
		lockdata = lockdata.. "???" .. client_name
		result = redis.call('setnx',key, lockdata)

		if result == 0 then

			local keydata = string.split(redis.call('get',key),'???')
			local time_out = tonumber(keydata[1])
			local data_client_name =  keydata[2]

			if cur_timestamp >  time_out then
				time_out = expire + cur_timestamp
				lockdata =  time_out .. "???" .. client_name

				if redis.call('setex',key,expire,lockdata)   then
					return    1

				end

				return 0
			end

			if data_client_name == client_name then

				return 1
			end

			return 0
		end

		if  redis.call('Expire',key,expire)  then
			return 1
		end

		return 0
`


func GetWLock() string{
	return wlock
}