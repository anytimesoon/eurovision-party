const base = "http://localhost:8080/"
const restricted = "restricted/api/"
const countryURL = base + restricted + "country/"
const userURL = base + restricted + "user/"
const voteURL = base + restricted + "vote/"

export const chatEP = "ws://localhost:8080/chat/connect";

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

export const auth = {
    LOGIN: base + "login",
    REGISTER: base + "register"
}