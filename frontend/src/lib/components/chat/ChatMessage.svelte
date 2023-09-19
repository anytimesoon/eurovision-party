<script lang="ts">
    import {CommentModel} from "$lib/models/classes/comment.model";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import {UserModel} from "$lib/models/classes/user.model";

    export let comment:CommentModel = new CommentModel()
    export let user:UserModel = new UserModel()
</script>

{#if $currentUser.id === user.id}
    <div class="flex w-full mt-2 space-x-3 max-w-xs ml-auto justify-end">
        <div>
            <p class="text-sm text-right">{user.name || ""}</p>
            <div class="bg-purple-800 text-white p-3 rounded-l-lg rounded-br-lg">

                <p class="text-sm">{comment.text}</p>
            </div>
            <span class="text-xs text-gray-500 leading-none">{comment.createdAt.getHours()}:{comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes()}</span>
        </div>
        <img class="flex-shrink-0 h-10 w-10 rounded-full" src={staticEP.IMG + user.icon} alt={user.name + "'s avatar"}>
    </div>
{:else if user.authLvl === authLvl.BOT}
    <div class="text-center mt-2 text-s">
        <p>{comment.text}</p>
    </div>
{:else}
    <div class="flex w-full mt-2 space-x-3 max-w-xs">
        <img class="flex-shrink-0 h-10 w-10 rounded-full" src={staticEP.IMG + user.icon} alt={user.name + "'s avatar"}>
        <div>
            <p class="text-sm">{user.name || ""}</p>
            <div class="bg-gray-300 p-3 rounded-r-lg rounded-bl-lg">

                <p class="text-sm">{comment.text}</p>
            </div>
            <span class="text-xs text-gray-500 leading-none">{comment.createdAt.getHours()}:{comment.createdAt.getMinutes() < 10 ? "0" + comment.createdAt.getMinutes() : comment.createdAt.getMinutes()}</span>
        </div>
    </div>
{/if}