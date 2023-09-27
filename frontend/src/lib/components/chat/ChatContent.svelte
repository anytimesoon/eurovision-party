<script lang="ts">
    import type {CommentModel} from "$lib/models/classes/comment.model";
    import type {UserModel} from "$lib/models/classes/user.model";
    import {currentUser} from "$lib/stores/user.store";

    export let comment:CommentModel
    export let user:UserModel

    function getHRTime(comment:CommentModel):string{
        return comment.createdAt.getHours() + ":" + (comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes())
    }

    $: textRight = $currentUser.id === user.id ? "text-right" : ""
    $: whiteText = $currentUser.id === user.id ? "text-white" : ""
    $: timeStyle = $currentUser.id === user.id ? "text-gray-300" : "text-right"
</script>

<span class="block pb-1 text-sm font-bold {textRight} {whiteText}">{user.name}</span>
<p class="{textRight} {whiteText} text-sm">{comment.text}</p>
<span class="text-[0.6rem] block pt-1 {timeStyle}">{getHRTime(comment)}</span>