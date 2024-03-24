<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import {onMount} from "svelte";
    import Spinner from "$lib/components/Spinner.svelte";
    import {socketRetryCount} from "$lib/stores/socketRetryCount.store";
    import {commentQueue} from "$lib/stores/commentQueue.store";

    export let comment:CommentModel
    export let isCurrentUser:boolean
    let commentElement:HTMLParagraphElement = document.createElement("p")

    $: if(comment) {
        const urlRegex = /((?:(http|https|Http|Https|rtsp|Rtsp):\/\/(?:(?:[a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(?:\%[a-fA-F0-9]{2})){1,64}(?:\:(?:[a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(?:\%[a-fA-F0-9]{2})){1,25})?\@)?)?((?:(?:[a-zA-Z0-9][a-zA-Z0-9\-]{0,64}\.)+(?:(?:aero|arpa|asia|a[cdefgilmnoqrstuwxz])|(?:biz|b[abdefghijmnorstvwyz])|(?:cat|com|coop|c[acdfghiklmnoruvxyz])|d[ejkmoz]|(?:edu|e[cegrstu])|f[ijkmor]|(?:gov|g[abdefghilmnpqrstuwy])|h[kmnrtu]|(?:info|int|i[delmnoqrst])|(?:jobs|j[emop])|k[eghimnrwyz]|l[abcikrstuvy]|(?:mil|mobi|museum|m[acdghklmnopqrstuvwxyz])|(?:name|net|n[acefgilopruz])|(?:org|om)|(?:pro|p[aefghklmnrstwy])|qa|r[eouw]|s[abcdeghijklmnortuvyz]|(?:tel|travel|t[cdfghjklmnoprtvwz])|u[agkmsyz]|v[aceginu]|w[fs]|y[etu]|z[amw]))|(?:(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9])\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[0-9])))(?:\:\d{1,5})?)(\/(?:(?:[a-zA-Z0-9\;\/\?\:\@\&\=\#\~\-\.\+\!\*\'\(\)\,\_])|(?:\%[a-fA-F0-9]{2}))*)?(?:\b|$)/gi;;
        const linkifiedText = comment.text.replace(urlRegex, function(url:string) {
            return '<a href="' + url + '">' + url + '</a>'
        })
        const parser = new DOMParser
        const htmlDoc = parser.parseFromString(linkifiedText, "text/html")
        commentElement.innerHTML = htmlDoc.body.innerHTML
    }

    function removeMessage() {
        commentQueue.removeMessage(comment.id)
    }

    $: contentTextStyle = isCurrentUser ? "text-typography-chat-me" : "text-typography-chat-you"
</script>

<div class="flex">
    <div>
        {#if comment.replyToComment}
            <div class="text-sm bg-canvas-secondary rounded px-3 py-1 border border-secondary">
                <span class="text-typography-chat-you">
                    {comment.replyToComment.text}
                </span>
            </div>
        {/if}

        <p bind:this={commentElement} class="{contentTextStyle} pr-8"></p>
    </div>


    <div class="flex flex-col-reverse flex-shrink">
        {#if comment.createdAt !== null}
            <span class="text-[0.6rem] block {contentTextStyle} text-right">
                {
                    comment.createdAt.getHours() + ":" +
                    (comment.createdAt.getMinutes() < 10 ?
                            "0" + comment.createdAt.getMinutes() :
                            comment.createdAt.getMinutes()
                    )
                }
            </span>
        {:else }
            <span class="text-[0.6rem] block {contentTextStyle} text-right">
                <Spinner size={"s"} thickness={"s"} color={"grey"}/>
            </span>

            {#if $socketRetryCount > 3}
                <span class="block p-0 m-0 -mt-2.5 text-warning cursor-pointer"
                      on:mouseup={removeMessage}
                      on:tap={removeMessage}>
                    ❌
                </span>
            {/if}
        {/if}


    </div>
</div>

