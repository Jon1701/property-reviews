-- Empty tables.
TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE management_companies CASCADE;
TRUNCATE TABLE properties CASCADE;

-- Create users.
INSERT INTO users(
	id_hash,
	email_address,
	password
)
VALUES (
	'user_9bf2f66e697e4b91b86789c899f900c1',
	'user@example.com',
	'$2a$10$UU3XrNJC2lbdHTcqlxld0u8GTYuhUgQLN454T6q.VsTMrfUc43oLW'
);

-- Create Property Management Companies.
INSERT INTO management_companies(id_hash, name, address_line1, address_line2, address_city, address_state, address_postal_code, address_country, website)
VALUES (
	'management_97ae61158ba747e19f0e9791a80f8028',
	'Enterprise Property Management',
	'1701 Kirk Avenue',
	'Suite 4834',
	'Toronto',
	'ON',
	'A3D7HY',
	'Canada',
	'https://www.example.com/enterprise-property-management'
),(
	'management_15af1325894d403e80e9f14f905c09cb',
	'Excelsior Property Management',
	'2000 Sulu Parkway',
	'Unit 17625',
	'Mississauga',
	'ON',
	'V3H8J7',
	'Canada',
	'https://www.example.com/excelsior-property-management'
),(
	'management_1f70063d453c4b56a4be712c4b255bcf',
	'Reliant Property Management',
	'1864 Terrell Road',
	NULL,
	'Etobicoke',
	'ON',
	'S5V8K8',
	'Canada',
	NULL
),(
	'management_24bcfe0e443841b78a074a26e4a9bab7',
	'Stargazer Property Management',
	'2893 Picard Lane',
	'Suite 001',
	'Markham',
	'ON',
	'T4V3S0',
	'Canada',
	'https://www.example.org/stargazer-property-management'
),(
	'management_df1c0b41427a48e89ecd9e7afce5d00d',
	'Defiant Property Management',
	'74205 Sisko Station',
	'Suite 9',
	'San Francisco',
	'CA',
	'90129',
	'United States of America',
	'https://www.example.org/defiant-pm'
),(
	'management_7d44f5f5ab39497c8626d32c5ea1d0bf',
	'Voyager Property Management',
	'74656 Janeway Way',
	'Block Delta',
	'North York',
	'ON',
	'J7B3X1',
	'Canada',
	'https://www.example.org/voypm'
),(
	'management_ad1bfd2a61564fc9bed052c0ef113a60',
	'NX Property Management',
	'01 Archer Avenue',
	NULL,
	'Broken Bow',
	'OK',
	'23682',
	'United States of America',
	'https://www.example.org/nx-property-management'
),(
	'management_9659fa3e55dd4c5a81c869f0a6806865',
	'Discovery Property Management',
	'1031 Lorca Road',
	NULL,
	'Mississauga',
	'ON',
	'X1T6J3',
	'Canada',
	'https://www.example.org/discovery-property-management'
),(
	'management_9e1b6d03ce3d4daf8ed85ff7878f4759',
	'Shenzou Property Management',
	'1227 Philippa Lane',
	NULL,
	'Toronto',
	'ON',
	'A3J8X1',
	'Canada',
	'https://www.example.org/shenzou-property-management'
),(
	'management_1484204e05e34acaa141a28d4aa53144',
	'Titan Property Management',
	'80102 Riker Avenue',
	NULL,
	'Toronto',
	'ON',
	'K9C491',
	'Canada',
	'https://www.example.org/titan-property-management'
),(
	'management_d5ab3213fed24f14a1446701e6b8c3e4',
	'Cerritos Property Management',
	'75567 Freeman Avenue',
	NULL,
	'Cerritos',
	'CA',
	'19254',
	'United States of America',
	'https://www.example.org/cerritos-property-management'
),(
	'management_b2e7316518a84c9f82dafb41b2d9f4df',
	'Grissom Property Management',
	'638 Esteban Avenue',
	NULL,
	'Etobicoke',
	'ON',
	'V2H8L6',
	'Canada',
	'https://www.example.org/grissom-property-management'
),(
	'management_e380dd6ee3ea47ccbe77ff4491657a44',
	'Bozeman Property Management',
	'1941 Bateson Avenue',
	NULL,
	'Markham',
	'ON',
	'B2D0K1',
	'Canada',
	'https://www.example.org/bozeman-property-management'
),(
	'management_c775dc6af9fe45f1b6b634974d06165a',
	'Rhode Island Property Management',
	'72701 Kim Boulevard',
	'Suite 2370',
	'Pickering',
	'ON',
	'J2B7A2',
	'Canada',
	'https://www.example.org/rhode-island-property-management'
),(
	'management_31b842ee728148e3885b730e52f4d72d',
	'Equinox Property Management',
	'72381 Ransom Avenue',
	'Suite 291572',
	'North York',
	'ON',
	'J2X1S1',
	'Canada',
	'https://www.example.org/equinox-property-management'
),(
	'management_9906ca990e2d4d59b40af15a00c7e2f7',
	'Franklin Property Management',
	'326 Edison Lane',
	'Suite 291572',
	'Hamilton',
	'ON',
	'N2C9E1',
	'Canada',
	'https://www.example.org/franklin-property-management'
),(
	'management_71974f49e20746bfaab7be3a9266b0c3',
	'Columbia Property Management',
	'02 Hernandez Road',
	'Suite 2811225',
	'Hamilton',
	'ON',
	'C3A0G2',
	'Canada',
	'https://www.example.org/columbia-property-management'
),(
	'management_5a3fccb945444c068dfe4772baa2a74c',
	'Kelvin Property Management',
	'0514 Robau Crossing',
	NULL,
	'London',
	'ON',
	'S7K9S5',
	'Canada',
	'https://www.example.org/kelvin-property-management'
),(
	'management_0e950de4af124c1bb42fa05fe766c489',
	'Leondegrance Property Management',
	'2176 Uhura Avenue',
	NULL,
	'Guelph',
	'ON',
	'H7S1H1',
	'Canada',
	'https://www.example.org/leondegrance-property-management'
),(
	'management_cc4d5f0d81864801b2b66837f7e954b5',
	'Lakota Property Management',
	'42768 Benteen Road',
	NULL,
	'Windsor',
	'ON',
	'W8I3Z1',
	'Canada',
	'https://www.example.org/lakota-property-management'
),(
	'management_c45da69d6e994678ae5ad60ced9a2448',
	'Armstrong Property Management',
	'317856 Imahara Parkway',
	NULL,
	'Simcoe',
	'ON',
	'Y2X8H8',
	'Canada',
	'https://www.example.org/armstrong-property-management'
),(
	'management_b6765cb40a6745b785aa3c5dccc08c54',
	'Vengeance Property Management',
	'31 Marcus Road',
	'Section 31',
	'Ottawa',
	'ON',
	'U7D3A9',
	'Canada',
	NULL
),(
	'management_43f09cabaa6d4baea9319c22861c23e0',
	'Charon Property Management',
	'1 Maddox Road',
	NULL,
	'Kingston',
	'ON',
	'C9T5W8',
	'Canada',
	NULL
),(
	'management_63b4756176b444238ba1f26ecb1a6c14',
	'Challenger Property Management',
	'71099 La Forge Avenue',
	'Unit 1725629',
	'Thunder B ay',
	'ON',
	'H2H18A',
	'Canada',
	'https://www.example.org/challenger-property-management'
);

-- Create Properties.
INSERT INTO properties(
	id_hash,
	management_company_id_hash,
	address_line1,
	address_line2,
	address_city,
	address_state,
	address_postal_code,
	address_country,
	property_type,
	building_type,
	neighborhood
)
VALUES (
	'property_a8f58132165b4d11b06dd19050294f0e',
	'management_df1c0b41427a48e89ecd9e7afce5d00d',
	'50000 Yonge Street',
	NULL,
	'Toronto',
	'ON',
	'I8H7D2',
	'Canada',
	'residential',
	'apartment',
	'Willowdale North'	
),(
	'property_15284db05752437a925bfbef0f7e0235',
	NULL,
	'287261 Yonge Street',
	NULL,
	'Toronto',
	'ON',
	'M173I1',
	'Canada',
	'residential',
	'apartment',
	'Willowdale North'	
),(
	'property_77c6433c3ebc40ad8e63f313719da2da',
	'management_5a3fccb945444c068dfe4772baa2a74c',
	'181218236 Yonge Street',
	'Unit 1622',
	'Toronto',
	'ON',
	'I8G1R1',
	'Canada',
	'residential',
	'detached',
	'Willowdale Norther North'	
),(
	'property_05617d0f8cec4c83891608bbca739f08',
	NULL,
	'381657 Yonge Street',
	'Suite 2815',
	'North York',
	'ON',
	'K8H6G1',
	'Canada',
	'residential',
	'detached',
	'South Willowdale'	
),(
	'property_17725d58b8224f8c89a6a07507851f11',
	NULL,
	'389228 Sheppard Avenue West',
	'Suite 2815',
	'North York',
	'ON',
	'C2T1X6',
	'Canada',
	'residential',
	'cottage',
	'Sheppard District C'	
),(
	'property_ad454a3fbb32485cbd198b4a6f29bd2a',
	'management_43f09cabaa6d4baea9319c22861c23e0',
	'3716252 Sheppard Avenue West',
	'Unit 2815',
	'North York',
	'ON',
	'K6F3A1',
	'Canada',
	'residential',
	'bungalow',
	'Sheppard District K'	
),(
	'property_146e57d887f1407ca173fb9385187720',
	'management_cc4d5f0d81864801b2b66837f7e954b5',
	'37162518 York Mills Avenue',
	'Unit 172561',
	'Toronto',
	'ON',
	'8J23Z1',
	'Canada',
	'residential',
	'basement-suite',
	'York Mills District V'	
),(
	'property_07675e0b885144b99ce7dee61ad9bb1a',
	'management_9659fa3e55dd4c5a81c869f0a6806865',
	'3591706 Davisville Avenue',
	'Unit 236981',
	'Toronto',
	'ON',
	'D1H2K1',
	'Canada',
	'residential',
	'townhome',
	'Davisville District V'	
),(
	'property_29cda3cd4ea14d2a8e36ba243d1ffe3c',
	NULL,
	'6272116 Finch Avenue',
	NULL,
	'North York',
	'ON',
	'J8D1K1',
	'Canada',
	'residential',
	'townhome',
	'Willowdale District ACS'	
),(
	'property_8eead691806b46e78d74f2893252040d',
	NULL,
	'5361851 Finch Avenue West',
	NULL,
	'North York',
	'ON',
	'A7H5F5',
	'Canada',
	'residential',
	'townhome',
	'Willowdale District KMM'	
),(
	'property_297c0fe35b65475eb8093066bc5d5d71',
	NULL,
	'233907163 Finch Avenue West',
	NULL,
	'North York',
	'ON',
	'L9K7H7',
	'Canada',
	'residential',
	'townhome',
	'Willowdale District NBM'	
),(
	'property_79fc38b6b25a47c3b3b9ce7b2e0dc47b',
	NULL,
	'362839061 Don Mills Road',
	NULL,
	'Toronto',
	'ON',
	'K8B3A1',
	'Canada',
	'residential',
	'townhome',
	'Fairview District AZR'	
),(
	'property_103b3402ae7b40cd85db60daea3c414a',
	NULL,
	'3262152 Sheppard Avenue East',
	NULL,
	'Toronto',
	'ON',
	'H6GB5A',
	'Canada',
	'residential',
	'townhome',
	'Sheppard District CZA'	
);
