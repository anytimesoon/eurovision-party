<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";

    export let comment:CommentModel
    export let isCurrentUser:boolean

    $: contentTextStyle = isCurrentUser ? "text-typography-chat-me pl-10" : "text-typography-chat-you pr-10"
    $: timeStyle = isCurrentUser ? "" : "text-right"
</script>


<p class="{contentTextStyle} text-sm">{comment.text}</p>
<span class="text-[0.6rem] -mt-2 block text-typography-grey {timeStyle}">
    {
        comment.createdAt.getHours() + ":" +
        (comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes())
    }
</span>