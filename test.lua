function readFile(file)
    local f = assert(io.open(file, "rb"))
    local content = f:read("*all")
    f:close()
    return content
end

Items = readFile("./testdata/items.json")

function init(args)
    math.randomseed(os.time() * os.clock() * 1000 + math.modf(io.popen("date +%S%N"):read("*a") / 1000000))
end

function request()
    local id = ""
    local characterSet = "abcdefghijklmnopqrstuvwxyz"
    local keyLength = 10

    for i = 1, keyLength do
        local rand = math.random(#characterSet)
        id = id .. string.sub(characterSet, rand, rand)
    end

    local body = '{"request_id":"'.. id ..'","items": ' .. Items .. '}'
    return wrk.format("POST", "/api/calc", {["Content-Type"] = "application/json"}, body)
end
