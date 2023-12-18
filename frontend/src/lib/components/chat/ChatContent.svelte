<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";

    export let comment:CommentModel
    export let isCurrentUser:boolean

    function linkify(text) {
        var urlRegex = /((?:(http|https|Http|Https|rtsp|Rtsp):\/\/(?:(?:[a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(?:\%[a-fA-F0-9]{2})){1,64}(?:\:(?:[a-zA-Z0-9\$\-\_\.\+\!\*\'\(\)\,\;\?\&\=]|(?:\%[a-fA-F0-9]{2})){1,25})?\@)?)?((?:(?:[a-zA-Z0-9][a-zA-Z0-9\-]{0,64}\.)+(?:(?:aero|arpa|asia|a[cdefgilmnoqrstuwxz])|(?:biz|b[abdefghijmnorstvwyz])|(?:cat|com|coop|c[acdfghiklmnoruvxyz])|d[ejkmoz]|(?:edu|e[cegrstu])|f[ijkmor]|(?:gov|g[abdefghilmnpqrstuwy])|h[kmnrtu]|(?:info|int|i[delmnoqrst])|(?:jobs|j[emop])|k[eghimnrwyz]|l[abcikrstuvy]|(?:mil|mobi|museum|m[acdghklmnopqrstuvwxyz])|(?:name|net|n[acefgilopruz])|(?:org|om)|(?:pro|p[aefghklmnrstwy])|qa|r[eouw]|s[abcdeghijklmnortuvyz]|(?:tel|travel|t[cdfghjklmnoprtvwz])|u[agkmsyz]|v[aceginu]|w[fs]|y[etu]|z[amw]))|(?:(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9])\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[1-9]|0)\.(?:25[0-5]|2[0-4][0-9]|[0-1][0-9]{2}|[1-9][0-9]|[0-9])))(?:\:\d{1,5})?)(\/(?:(?:[a-zA-Z0-9\;\/\?\:\@\&\=\#\~\-\.\+\!\*\'\(\)\,\_])|(?:\%[a-fA-F0-9]{2}))*)?(?:\b|$)/gi;;
        return text.replace(urlRegex, function(url) {
            return '<a href="' + url + '">' + url + '</a>';
        });
    }

    $: contentTextStyle = isCurrentUser ? "text-typography-chat-me" : "text-typography-chat-you"
</script>

<div class="flex">
    <div>
        {#if comment.replyToComment}
            <div class="text-sm bg-chat-bubble-you rounded px-3 py-1">
                <span class="text-typography-chat-you">{comment.replyToComment.text}</span>
            </div>

        {/if}

        <p class="{contentTextStyle} pr-8">{linkify(comment.text)}</p>
    </div>


    <div class="flex flex-col-reverse flex-shrink">
        <span class="text-[0.6rem] block {contentTextStyle} text-right">
    {
        comment.createdAt.getHours() + ":" +
        (comment.createdAt.getMinutes() < 10 ?
                "0" + comment.createdAt.getMinutes() :
                comment.createdAt.getMinutes()
        )
    }
</span>
    </div>
</div>

