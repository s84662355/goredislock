package script


var unwlock string = `

    function string.split(input, delimiter)
        input = tostring(input)
        delimiter = tostring(delimiter)
        if (delimiter=='') then return false end
        local pos,arr = 0, {}
        -- for each divider found
        for st,sp in function() return string.find(input, delimiter, pos, true) end do
            table.insert(arr, string.sub(input, pos, st - 1))
            pos = sp + 1
        end
        table.insert(arr, string.sub(input, pos))
        return arr
    end

    local key   = KEYS[1]
    local result = redis.call('get',key)
    local client_name =  ARGV[1]
    if result == nil   then
       if redis.call('DEL',key) then
            return  1
       end
       return 0
    end
    local keydata = string.split(result ,'???')
    if client_name  == keydata[2] and  redis.call('DEL',key)  then
          return 1
     end

    return 0


`


func GetUnWLock() string{
	return unwlock
}