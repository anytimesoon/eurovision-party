const apiBase = "http://localhost:8080/"
const feBase = "http://localhost:5173/"
const restricted = "restricted/"
const svelteAPI = "endpoints/"
const goAPI = "api/"
const countrySvelteURL = feBase + svelteAPI + "country/"
const userSvelteURL = feBase + svelteAPI + "user/"
const voteSvelteURL = feBase + svelteAPI + "vote/"
const countryGoURL = apiBase + restricted + goAPI + "country/"
const userGoURL = apiBase + restricted + goAPI + "user/"
const voteGoURL = apiBase + restricted  + goAPI + "vote/"



export const chatEP = "ws://localhost:8080/restricted/chat/connect";

export const countrySvelteEP = {
    ALL: countrySvelteURL,
    UPDATE: countrySvelteURL,
    FIND_ONE: countrySvelteURL,
    PARTICIPATING: countrySvelteURL + "participating"
}

export const countryGoEP = {
    ALL: countryGoURL,
    UPDATE: countryGoURL,
    FIND_ONE: countryGoURL,
    PARTICIPATING: countryGoURL + "participating"
}

export const staticEP = {
    IMG: apiBase
}

export const userSvelteEP = {
    ALL: userSvelteURL,
    UPDATE: userSvelteURL,
    FIND_ONE: userSvelteURL,
    REMOVE: userSvelteURL,
    REGISTERED: userSvelteURL + "registered"
}

export const userGoEP = {
    ALL: userGoURL,
    UPDATE: userGoURL,
    FIND_ONE: userGoURL, //append with user slug
    REMOVE: userGoURL,
    REGISTERED: userGoURL + "registered"
}

export const voteSvelteEP = {
    CREATE: voteSvelteURL,
    UPDATE: voteSvelteURL,
    BY_COUNTRY_AND_USER: voteSvelteURL + "countryanduser/" //append with country slug
}

export const voteGoEP = {
    CREATE: voteGoURL,
    UPDATE: voteGoURL,
    BY_COUNTRY_AND_USER: voteGoURL + "countryanduser/" //append with country slug
}

export const authEP = {
    LOGIN: apiBase + "login",
    REGISTER: userSvelteURL + "register"
}