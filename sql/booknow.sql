
/*********** BookNow Api master data ***********/

/*Types*/
INSERT INTO public.type
	(id, description)
	VALUES 
	(1, 'Contact'),
	(2, 'Social'),
	(3, 'Person'),
	(4, 'Message'),
	(5, 'Document'),
	(6, 'Folder')

/*References*/
INSERT INTO public.reference
	(id, description, type_id)
	VALUES 
	(1, 'Email Address', 1),
	(2, 'Mobile', 1),
	(3, 'Residential Address', 1),
	(4, 'Website Address', 1),
	(5, 'Social Address', 1),
	(6, 'Facebook', 2),
	(7, 'LinkedIn', 2),
	(8, 'GitHub', 2),
	(9, 'Instagram', 2),
	(10, 'Whatsapp', 2),
	(11, 'Landline', 1),
	(12, 'Anonymous', 3),
	(13, 'User', 3),
	(14, 'Client', 3),
	(15, 'Tenant', 3),
	(16, 'Subcontractor', 3),
	(17, 'Partner', 3),
	(18, 'Incident', 4),
	(19, 'Complaint', 4),
	(20, 'Enquiry', 4),
	(21, 'Suggestion', 4),
	(22, 'Introduction', 4),
	(23, 'Owner', 3)
	(24, 'Police Check', 5),
	(25, 'Working with children check', 5),
	(26, 'Visa', 5),
	(27, 'Student Id', 5),
	(28, 'Resume', 5),
	(29, 'Passport first page', 5),
	(30, 'Passport last page', 5)
	(31, 'ABN', 1),
	(32, 'Documents', 6)

