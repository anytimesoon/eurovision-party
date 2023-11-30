<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";

    export let comment:CommentModel
    export let isCurrentUser:boolean

    $: contentTextStyle = isCurrentUser ? "text-typography-chat-me" : "text-typography-chat-you"
    $: contentPadding = isCurrentUser ? "pl-10" : "pr-10"
    $: timeStyle = isCurrentUser ? "" : "text-right"
</script>


<p class="{contentTextStyle} {contentPadding} text-sm">{comment.text}</p>
<span class="text-[0.6rem] -mt-2 block {contentTextStyle} {timeStyle}">
    {
        comment.createdAt.getHours() + ":" +
        (comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes())
    }
</span>