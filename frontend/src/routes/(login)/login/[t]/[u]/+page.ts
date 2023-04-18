// eslint-disable-next-line @typescript-eslint/no-unused-vars,@typescript-eslint/ban-ts-comment
// @ts-ignore
/** @type {import('./$types').PageLoad} */
export function load({ params }) {
    return {
        t: params.t,
        u: params.u
    }
}