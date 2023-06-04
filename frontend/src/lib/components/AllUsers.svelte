<script lang="ts">
    import {staticEP} from "$lib/models/enums/endpoints.enum.js";
</script>

<ul id="full-list" style="text-decoration: none;">
        {#each users as [id, user]}
            <li>
                <img src="{staticEP.IMG + user.icon}" alt="avatar missing" width="50">
                {user.name}
            </li>
        {/each}
</ul>