import pandas as pd
import faiss
from sentence_transformers import SentenceTransformer, util
import sys

def runn(test_course_code, test_description):
    model = SentenceTransformer('all-MiniLM-L6-v2')
    catalog_file = "/mnt/Disk_2/hack_cbs/project/Credis/model/course-catalog.csv"
    catalog_df = pd.read_csv(catalog_file)
    catalog_df['Description'] = catalog_df['Description'].fillna("")
    descriptions = catalog_df['Description'].tolist()
    description_embeddings = model.encode(descriptions, convert_to_tensor=False)
    dimension = description_embeddings[0].shape[0]
    index = faiss.IndexFlatL2(dimension)
    index.add(description_embeddings)

    def find_similar_courses(course_code, description, catalog_df, threshold=0.4):
        query_embedding = model.encode([description], convert_to_tensor=False)
        k = len(descriptions)
        distances, indices = index.search(query_embedding, k)
        unique_matches = {}
        for distance, idx in zip(distances[0], indices[0]):
            similarity_score = 1 / (1 + distance)
            if similarity_score > threshold:
                matched_course_code = f"{catalog_df.iloc[idx]['Subject']} {catalog_df.iloc[idx]['Number']}"
                matched_description = catalog_df.iloc[idx]['Description']
                if matched_course_code not in unique_matches or similarity_score > unique_matches[matched_course_code][2]:
                    unique_matches[matched_course_code] = (matched_course_code, matched_description, similarity_score)
        matches = sorted(unique_matches.values(), key=lambda x: x[2], reverse=True)
        return [(code, desc) for code, desc, _ in matches]

    matched_courses = find_similar_courses(test_course_code, test_description, catalog_df, threshold=0.5)
    file = open("output.txt","w")
    if matched_courses:
        file.write("Courses Found!\n")
        for course in matched_courses:
            print(f"Course Code: {course[0]}")
            print(f"Description: {course[1]}")
            file.write(course[0])
            file.write("\n")
    else:
        print("No matches found.")
        file.write("No match")

runn(sys.argv[1], sys.argv[2])
