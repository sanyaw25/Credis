# Credis
**Problem**-
For admissions, lateral entries and online courses such as NPTEL, universities often perform manual reviews doing course-similarity search causing time delays. Course names are not the same at every institution (digital design is digital logic design, linear algebra is linear algebra & differential equations ). In India, no such solution has been addressed even by ABC.

**Solution**-
We propose a secure AI-powered credit transfering system for students which proposes authentication for students, their own dashboard where they can upload their academic transcripts, with attestations for maintaining track and history. Institutions can easily review courses seeing accurate similar matching courses using advanced ML techniques based on course syllabus, course names, description, credit hours on university catalogs and approve credit and student transfers thus automating the once laborious process.

**Features**
1. Attestations and Https implementation
2. AI-powered Similarity course search to reduce manual overview
3.Auth0 & MongoDB for authentication & data record storage	
4.Easy to understand and use, just need to build account and upload your transcript.

**Working**
1. Account creation using auth0 for students
2. After login, dashboard contains profile and option to upload academic transcripts a
3. PDF text extraction using PyMuPDF library and NLP model which further encodes into vectors (the model had a smaller footprint than BERT).
4. Attestations are also added with each upload to keep easy track of history.
5. Model will compare it with a csv file containing a course catalog and display matching courses that can be transferred to any institution

**Credits**
Sanya Wadhawan
Ayushman Agrawal Hingorani

