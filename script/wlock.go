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
		local lockdata =  expire + cur_timestamp
		lockdata = lockdata.. "???" .. client_name
		 

		return  redis.call('set',key, lockdata,'nx','ex',expire)

`


func GetWLock() string{
	return wlock
}