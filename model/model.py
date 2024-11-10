import pandas as pd
import faiss
from sentence_transformers import SentenceTransformer, util
import sys

def runn(test_course_code, test_description):
    # Load the sentence-transformer model for embeddings
    model = SentenceTransformer('all-MiniLM-L6-v2')  # A small, fast transformer model

    # Load catalog data from "course-catalog.csv" into a DataFrame
    catalog_file = "/mnt/Disk_2/hack_cbs/project/Credis/model/course-catalog.csv"
    catalog_df = pd.read_csv(catalog_file)

    # Handle missing descriptions by filling NaN values with an empty string
    catalog_df['Description'] = catalog_df['Description'].fillna("")

    # Generate embeddings for each course description
    descriptions = catalog_df['Description'].tolist()
    description_embeddings = model.encode(descriptions, convert_to_tensor=False)

    # Build FAISS index
    dimension = description_embeddings[0].shape[0]  # Dimensionality of the embeddings
    index = faiss.IndexFlatL2(dimension)  # Using L2 distance for similarity search

    # Add embeddings to the FAISS index
    index.add(description_embeddings)

    # Define a function to find unique similar courses using FAISS
    def find_similar_courses(course_code, description, catalog_df, threshold=0.4):
        # Encode the input description to get its embedding
        query_embedding = model.encode([description], convert_to_tensor=False)

        # Search in FAISS for similar course descriptions
        k = len(descriptions)  # Retrieve all possible matches
        distances, indices = index.search(query_embedding, k)

        # Dictionary to store unique matches by course code with the highest similarity score
        unique_matches = {}
        for distance, idx in zip(distances[0], indices[0]):
            # Convert FAISS distance to a similarity score (inverse of distance)
            similarity_score = 1 / (1 + distance)
            if similarity_score > threshold:
                matched_course_code = f"{catalog_df.iloc[idx]['Subject']} {catalog_df.iloc[idx]['Number']}"
                matched_description = catalog_df.iloc[idx]['Description']

                # If the course code is new or has a higher similarity score, add/update it in unique_matches
                if matched_course_code not in unique_matches or similarity_score > unique_matches[matched_course_code][2]:
                    unique_matches[matched_course_code] = (matched_course_code, matched_description, similarity_score)

        # Convert the unique_matches dictionary to a sorted list by similarity score
        matches = sorted(unique_matches.values(), key=lambda x: x[2], reverse=True)
        return [(code, desc) for code, desc, _ in matches]

    # Example usage: Find similar courses
    matched_courses = find_similar_courses(test_course_code, test_description, catalog_df, threshold=0.5)
    file = open("output.txt","w")
    # Display matched courses
    if matched_courses:
        file.write("Courses Found!")
        for course in matched_courses:
            print(f"Course Code: {course[0]}")
            print(f"Description: {course[1]}")
            file.write(course[0])
            file.write("\n")
            
    else:
        print("No matches found.")
        file.write("No match")


runn(sys.argv[1], sys.argv[2])