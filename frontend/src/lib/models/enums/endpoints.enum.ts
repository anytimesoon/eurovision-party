const apiBase = "http://localhost:8080/"
const feBase = "http://localhost:3000/"
const restricted = "restricted/api/"
const countryURL = apiBase + restricted + "country/"
const userURL = apiBase + restricted + "user/"
const voteURL = apiBase + restricted + "vote/"
//TODO: tidy up endpoints

// Backend endpoints
export const chatEP = "ws://localhost:8080/chat/connect";

export const countryEP = {
    ALL: countryURL,
    UPDATE: countryURL,
    FIND_ONE: countryURL
}

export const staticEP = {
    IMG: apiBase
}

export const userEP = {
    ALL: userURL,
    UPDATE: userURL,
    FIND_ONE: userURL,
    REMOVE: userURL,
    REGISTERED: userURL + "registered"
}

export const voteEP = {
    CREATE: voteURL,
    UPDATE: voteURL
}

export const authEP = {
    LOGIN: apiBase + "login",
    REGISTER: userURL + "register",
    FE_LOGIN: feBase + "login/"
}

// Frontend enpoints

export const userFeEP = {
    SINGLE_USER: feBase + "person"
}