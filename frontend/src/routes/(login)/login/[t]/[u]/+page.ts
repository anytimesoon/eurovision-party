// eslint-disable-next-line @typescript-eslint/no-unused-vars,@typescript-eslint/ban-ts-comment
// @ts-ignore
import {sendCreateOrUpdate, sendGet} from "$lib/helpers/sender.helper";

/** @type {import('./$types').PageLoad} */
// @ts-ignore
export async function load({ params }) {
    return {
        token: params.t,
        userId: params.u
    }
}