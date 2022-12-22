const base = "http://localhost:8080/"
const restricted = "restricted/"
const api = "api/"
const countryURL = base + restricted + api + "country/"
const userURL = base + restricted + api + "user/"
const voteURL = base + restricted + api + "vote/"
const chatURL = base + restricted + "chat/"

export const countryEP = {
    ALL: countryURL,
    UPDATE: countryURL,
    FIND_ONE: countryURL
}

export const userEP = {
    ALL: userURL,
    UPDATE: userURL,
    FIND_ONE: userURL,
    REMOVE: userURL
}

export const voteEP = {
    CREATE: voteURL,
    UPDATE: voteURL
}

export const chatEP = {
    CONNECT: chatURL + "connect/"
}

export const auth = {
    LOGIN: base + "login",
    REGISTER: base + "register"
}