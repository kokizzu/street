<script>
	/** @typedef {import('../_types/property').Property} TypeProperty} */
	/** @typedef {import('../_types/property').PropertyUS} TypePropertyUS} */
	/** @typedef {import('../_types/property').PropertyWithNote} PropertyWithNote} */
	/** @typedef {import('../_types/master').Field} Field*/

	import { Icon } from '../node_modules/svelte-icons-pack/dist';
	import {
		LuBath, LuBed, LuArmchair, LuSquareDashedBottom
	} from '../node_modules/svelte-icons-pack/dist/lu';
	import {
		RiMediaImageLine, RiBuildingsHome3Line, RiDesignPencilLine,
		RiMapMapPinLine
	} from '../node_modules/svelte-icons-pack/dist/ri';
	import PillBox from './PillBox.svelte';
	import {
		localeDatetime, M2ToPing, getApprovalState, formatPrice
	} from './formatter';
	import { onMount } from 'svelte';
	
	export let property = /** @type {TypeProperty | TypePropertyUS | PropertyWithNote} */ ({});
	export let meta = /** @type {Field[]} */ ([]);
	let approvalStatus = /** @type {string} */ ('approved');
	let noteObj = /** @type {Record<string, any>} */ ({});

	onMount(() => {
		try {
			noteObj = JSON.parse( property.note );
		} catch (/** @type {any | Error} */ e) {
			noteObj = {
				contactPhone: '',
				contactEmail: '',
				about: property.note
			}
		}
		approvalStatus = getApprovalState( property.approvalState );
		if (property.countryCode !== 'US') {
			property.sizeM2 = String(M2ToPing(parseInt(property.sizeM2)));
		}
	});
</script>
  
<div class="property-note">
	<div class="property-main">
		<div class="images">
			{#if property.images && property.images.length}
				<img src={property.images[0]} alt=""/>
			{:else}
				<div class="img-empty">
					<Icon
						src={RiMediaImageLine}
						size="60"
						color="var(--gray-006)"
					/>
				</div>
			{/if}
		</div>
		<div class="info">
			<div class="col-1">
				<div class="left">
						<div class={(property.purpose).includes('sale')
							? 'purpose sale'
							: 'purpose rent'
						}>
							{(property.purpose).includes('sale')
								? 'For Sale'
								: 'For Rent'
							}
						</div>
						<div class="property-type">
							<Icon
								size="15"
								src={RiBuildingsHome3Line}
							/>
							<span>{property.houseType || 'House'}</span>
						</div>
						<div class="approve-status {approvalStatus}">
							{approvalStatus}
						</div>
				</div>
				<a class="btn-edit" href="/realtor/propertyOld/{property.id}">
					<Icon
						color="#FFF"
						size="17"
						src={RiDesignPencilLine}/>
					<span>Edit Property</span>
				</a>
			</div>
			<div class="col-2">
				<div class="price">
					<h1>
						{formatPrice(parseInt(property.lastPrice), 'USD')}
					</h1>
				</div>
				<span class="agency-fee">Agency Fee {property.agencyFeePercent}%</span>
				<div class="address">
					<Icon
						size="18"
						src={RiMapMapPinLine}
						color="var(--orange-006)"
					/>
					<span>
						{property.formattedAddress || property.address}
					</span>
				</div>
			</div>
		</div>
	</div>
	<div class="property-secondary">
		<div class="features">
			<div class="item">
				<b>{property.bedroom}</b>
				<div class="labels">
					<Icon
						src={LuBed}
						size="15"
					/>
					<span>Bedroom</span>
				</div>
			</div>
			<div class="item">
				<b>{property.bathroom}</b>
				<div class="labels">
					<Icon
						src={LuBath}
						size="15"
					/>
					<span>Bathroom</span>
				</div>
			</div>
			<div class="item">
				<b>{property.bathroom}</b>
				<div class="labels">
					<Icon
						src={LuArmchair}
						size="15"
					/>
					<span>Living Room</span>
				</div>
			</div>
			<div class="item">
				<b>{property.sizeM2}</b>
				<div class="labels">
					<Icon
						src={LuSquareDashedBottom}
						size="14"
					/>
					<span>
						{ property.countryCode === 'US' ? 'm2' : 'ping' }
					</span>
				</div>
			</div>
		</div>
	</div>
	<div class="property_attributes">
		{#each meta as m}
			{#if property[ m.name ]}
				{#if m.inputType === 'datetime'}
					<PillBox
						label={m.label}
						content={localeDatetime(property[m.name])}
					/>
				{:else}
					{#if m.name === 'note' && ( noteObj && noteObj.about)}
						<PillBox
							label={m.label}
							content={noteObj.about}
						/>
					{:else}
						<PillBox
							label={m.label}
							content={property[m.name]}
						/>
					{/if}
				{/if}
			{/if}
		{/each}
	</div>
</div>
  
<style>
	.unit_toggle {
		border: none;
		background: transparent;
		position: relative;
		cursor: pointer;
	}
  
	.property-note {
		height: fit-content;
		background-color: var(--gray-001);
		border-radius: 12px;
		padding: 20px;
		display: flex;
		flex-direction: column;
		border: 1px solid var(--gray-002);
		gap: 20px;
		position: relative;
	}
	
	.property-main {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;
		align-content: flex-start;
		justify-content: flex-start;
		align-items: flex-start;
		gap: 15px;
	}
  
	.property-main .images {
		width: 330px;
		height: 170px;
		flex: none;
		overflow: hidden;
		border-radius: 8px;
		border: 1px solid var(--gray-003);
	}
  
	.property-main .images img {
		object-fit: cover;
		width: 100%;
		height: 100%;
	}
  
	.property-main .images .img-empty {
		border-radius: 8px;
		object-fit: cover;
		width: 100%;
		height: 100%;
		background-color: var(--gray-001);
		display: flex;
		justify-content: center;
		align-items: center;
	}
  
	.property-main .info {
		flex-grow: 1;
		display: flex;
		flex-direction: column;
		gap: 10px;
	}
  
	.property-main .info .col-1 {
		display: flex;
		justify-content: space-between;
		align-items: center;
		flex-direction: row;
	}
  
	.property-main .info .col-1 .left {
		display: flex;
		flex-direction: row;
		gap: 10px;
		align-items: center;
		font-weight: 600;
	}

	.property-main .info .col-1 .left .approve-status {
		padding: 7px 18px;
		border-radius: 999px;
		text-transform: capitalize;
		text-decoration: none;
	}
	
	.property-main .info .col-1 .left .approve-status.approved {
		background-color: var(--green-transparent);
		color: var(--green-006);
		border: 1px solid var(--green-006);
	}
	
	.property-main .info .col-1 .left .approve-status.pending {
		background-color: var(--yellow-transparent);
		color: var(--yellow-006);
		border: 1px solid var(--yellow-006);
	}
	
	.property-main .info .col-1 .left .approve-status.rejected {
		background-color: var(--red-transparent);
		color: var(--red-006);
		border: 1px solid var(--red-006);
	}

	.property-main .info .col-1 .left .purpose,
	.property-main .info .col-1 .left .property-type {
		padding: 7px 18px;
		border-radius: 999px;
		text-transform: capitalize;
		text-decoration: none;
		background-color: var(--orange-transparent);
		color: var(--orange-006);
		border: 1px solid var(--orange-006);
		display: flex;
		flex-direction: row;
		gap: 5px;
		align-items: center;
	}

	.property-main .info .col-1 .btn-edit {
		background-color: var(--orange-006);
		color: #FFF;
		padding: 10px 15px;
		border-radius: 999px;
		display: flex;
		flex-direction: row;
		gap: 10px;
		text-decoration: none;
	}

	.property-main .info .col-1 .btn-edit:hover {
		background-color : var(--orange-005);
	}
  
	.property-main .info .col-2 {
		display: flex;
		flex-direction: column;
		gap: 10px;
	}

	.property-main .info .col-2 .price {
		display: flex;
		flex-direction: row;
		gap: 10px;
		align-items: center;
	}
  
	.property-main .info .col-2 .price h1 {
		margin: 0;
		font-size: 35px;
	}
  
	.property-main .info .col-2 .agency-fee {
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 5px 13px;
		border-radius: 5px;
		border: 1px solid var(--gray-003);
		font-size: 14px;
		width: fit-content;
	}
  
	.property-main .info .col-2 .address span {
		flex-shrink: 1;
	}

	.property-main .info .col-2 .address {
		display: flex;
		flex-direction: row;
		gap: 10px;
		align-items: center;
	}
  
	.property-secondary {
		display: flex;
		flex-direction: column;
		gap: 15px;
	}

	.property-secondary .features {
		display: flex;
		flex-direction: row;
		justify-content: flex-start;
		align-items: center;
		gap: 15px;
	}

	.property-secondary .features .item {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: flex-end;
		gap: 5px;
	}

	.property-secondary .features .item b {
		font-size: 35px;
	}

	.property-secondary .features .item .labels {
		display: flex;
		flex-direction: row;
		gap: 5px;
		align-items: center;
		margin-bottom: 8px;
	}
  </style>