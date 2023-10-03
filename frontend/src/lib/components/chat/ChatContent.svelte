<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import type {UserModel} from "$lib/models/classes/user.model";

    export let comment:CommentModel
    export let user:UserModel
    export let isCurrentUser:boolean


    $: userNameStyle = isCurrentUser ? "text-right" : ""
    $: contentTextStyle = isCurrentUser ? "text-typography-chat-me" : "text-typography-chat-you"
    $: timeStyle = isCurrentUser ? "" : "text-right"
</script>

<p class="block pb-1 {userNameStyle}">{user.name}</p>
<p class="{contentTextStyle} text-sm">{comment.text}</p>
<span class="text-[0.6rem] block pt-1 text-typography-grey {timeStyle}">{comment.createdAt.getHours() + ":" + (comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes())}</span>